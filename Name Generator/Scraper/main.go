// Scrapes hobby data and saves it as a text file.

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// --------------------------------------------------------------------------------------------------------------------

// Const presets to change scraping options. Saved here for reference.

/*
	Hobby Data:
	const url=  "https://www.burlapandblue.com/list-of-hobbies/""
	const save_path =  "../hobbies.txt"
	const tag = "ul li strong"
	const option = "hobbies"

	State data:
	const url = "https://informationngr.com/us-zip-codes-list/"
	const save_path = "../states.txt"
	const tag = "tbody tr td"
	const option = "states"

	Job titles:
	const url = "https://zety.com/blog/job-titles"
	const save_path = "../job_titles.txt"
	const tag = "ul li"
	const option = "job_titles"
*/

// --------------------------------------------------------------------------------------------------------------------

func main() {

	// Uncomment / Comment to change scraped data.
	// const url = "https://informationngr.com/us-zip-codes-list/"
	// const save_path = "../states.txt"
	// const tag = "tbody tr td"
	// const option = "states"

	// const url = "https://www.burlapandblue.com/list-of-hobbies/"
	// const save_path = "../hobbies.txt"
	// const tag = "ul li strong"
	// const option = "hobbies"

	const url = "https://www.firstcareers.co.uk/job-titles/"
	const save_path = "../job_titles.txt"
	const tag = "div ul.clearfix li a"
	const option = "job titles"

	// -----------

	// Scrape data
	completed_data := scrape_data(tag, url, option)
	if len(completed_data) == 0 {
		log.Fatal(fmt.Errorf("scrape_data: nothing found\nURL:%v\nTag:%v", url, tag))
	}

	// -----------

	// Save data
	err := save_data(completed_data, save_path)
	if err != nil {
		fmt.Println("save data:", err)
	}
}

// --------------------------------------------------------------------------------------------------------------------

func scrape_data(tag string, url string, option string) []string {
	// Scrapes data found under a HTML from a URL.

	// -----------

	// Create a collector
	collector := colly.NewCollector()
	extensions.RandomUserAgent(collector)

	// -----------

	// Look for tag in html and save data.
	var raw_data string
	collector.OnHTML(tag, func(data *colly.HTMLElement) {
		raw_data += fmt.Sprintf("%v\n", data.Text)
	})
	collector.Visit(url)

	// -----------

	// Store scraped data
	completed_data := parse_data(raw_data, option)

	return completed_data

}

// --------------------------------------------------------------------------------------------------------------------

func parse_data(data string, option string) []string {
	// Turn scraped data into a slice of string data.

	// -----------

	// Slice incoming data based upon newlines.
	var completed_data []string
	scraped_data := strings.Split(data, "\n")

	// Parse data depending on type.
	switch option {
	case "hobbies":
		completed_data = hobbies(scraped_data)
	case "states":
		completed_data = states_zipcodes(scraped_data)
	case "job titles":
		completed_data = scraped_data
	}

	// -----------

	return completed_data
}

// --------------------------------------------------------------------------------------------------------------------

func save_data(data []string, path string) error {
	// Writes data to a textfile

	// -----------

	// Create file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// -----------

	// Write data to the file
	for _, value := range data {
		_, err := file.WriteString(fmt.Sprintf("%v\n", value))
		if err != nil {
			return err
		}
	}

	// -----------

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func hobbies(scraped_data []string) []string {
	// Takes scraped hobby data, parses it into a usable format for saving and removes maleformed information.

	// -----------

	var completed_data []string

	// Get non garbage data and add it to our dataset to save to a text file.
	for i := 0; i < len(scraped_data); i++ {

		// -----------

		// Ignore anything with dashes as go cannot remove them from the base data (reason unknown)
		if (strings.Contains(scraped_data[i], "-")) || (strings.Contains(scraped_data[i], " –")) {
			continue
		}

		// Split current line into its parts.
		element := strings.Split(scraped_data[i], " ")

		// Ignore weird data
		if len(element) <= 3 && len(element) > 0 {

			// Remove whitespace to make entry consistent
			for j := range element {
				element[j] = strings.TrimSpace(element[j])
			}

			// Rejoin into a string with spaces between each element, add to completed data.
			completed_data = append(completed_data, strings.Join(element, " "))
		}

		// -----------
	}

	// -----------

	return completed_data
}

// --------------------------------------------------------------------------------------------------------------------

func states_zipcodes(scraped_data []string) []string {
	// Takes scraped state and zipcode data and strips out needless information.

	// -----------

	var completed_data []string
	var data_points_used = 0

	// Get state data and zipcodes from the dataset. Everything past 100 entries is garbage data.
	for data_points_used < 50 {

		// Slice off the first 2 data points (State + Zipcode)
		current_data := scraped_data[:2]
		scraped_data = scraped_data[2:]

		// -----------

		// Create a data point to add to the completed data slice.
		state_data := strings.Split(current_data[0], " ")
		var state_name_data []string

		// Create the string of the state name using a loop until it meets unneeded data.
		i := 0
		for state_data[i] != "Zip" {
			state_name_data = append(state_name_data, state_data[i])
			i++
		}

		// Format the data.
		state_name := strings.Join(state_name_data, " ")
		state_slice := fmt.Sprintf("%v %v", state_name, current_data[1])

		// -----------

		// Add the completed slice data to the data we want to return
		completed_data = append(completed_data, state_slice)

		// -----------

		data_points_used++
	}

	// -----------

	return completed_data
}

// --------------------------------------------------------------------------------------------------------------------
