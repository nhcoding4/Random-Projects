// Functions used by go routines to collect, edit and parse raw data into something useful.

package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

// --------------------------------------------------------------------------------------------------------------------

func social_security(ch chan string) {
	// Generate a random social security number

	// -----------

	// Format of the social security number we want to return
	social_security := "xxx-xx-xxxx"

	// Replace the numbers with random digits
	for i := 0; i < 9; i++ {
		social_security = strings.Replace(social_security, "x", fmt.Sprintf("%v", rand.IntN(9)), 1)
	}

	// -----------

	ch <- social_security
	close(ch)
}

// --------------------------------------------------------------------------------------------------------------------

func name(ch chan string, errch chan error) {
	// Randomly selects a name and sex

	// -----------

	// Get a random male or female name

	const male_names = "../data/male_names.txt"
	const female_names = "../data/female_names.txt"

	// -----------

	// 0 = Male, 1 = Female
	number := rand.IntN(2)
	var path string

	if number == 0 {
		ch <- "male"
		path = male_names
	} else {
		ch <- "female"
		path = female_names
	}

	// -----------

	// Randomly select a first and last name
	func(path string) {

		data, err := load_data(path)
		if err != nil {
			errch <- fmt.Errorf("name: %v", err)
		}

		length := len(data)
		var previous_name string
		names_selected := 0

		// Ensure we dont select the same names.
		for names_selected < 2 {

			selection := data[rand.IntN(length)]

			if selection == previous_name {
				continue
			}

			ch <- selection
			names_selected++
		}

	}(path)

	// -----------

	close(ch)
}

// --------------------------------------------------------------------------------------------------------------------

func city(ch chan string, errch chan error, numch chan int) {
	// Generate a random city + zipcode pair

	// -----------

	// Load data from text file into memory.
	const states_file = "../data/states.txt"

	state_data, err := load_data(states_file)
	if err != nil {
		errch <- fmt.Errorf("city: loading data %v", err)
		return
	}

	// -----------

	// Turn data into something usable.
	data := make(map[string][]int)
	var keys []string

	for len(state_data) > 0 {

		// Get the first state/zipcode entry
		current_state := state_data[:1]
		state_data = state_data[1:]

		// Split the data into its parts
		split_data := strings.Split(current_state[0], " ")
		split_state := split_data[:len(split_data)-1]

		// -----------

		// Get the zip codes from the raw string
		zip_code_raw := split_data[len(split_data)-1:]

		for _, code := range zip_code_raw {

			// Get the range zip codes for a city appear in
			var lower_bound string
			var upper_bound string

			// Skip non numerical characters in the middle of each post code
			for j, char := range code {
				if j <= 4 {
					lower_bound += string(char)
				}
				if j > 7 {
					upper_bound += string(char)
				}
			}

			// -----------

			// Attempt to convert the string numbers into integers
			range_1, err := strconv.Atoi(strings.TrimSpace(lower_bound))
			if err != nil {
				errch <- fmt.Errorf("city: %v", err)
				return
			}
			range_2, err := strconv.Atoi(strings.TrimSpace(upper_bound))
			if err != nil {
				errch <- fmt.Errorf("city: %v", err)
				return
			}

			// -----------

			// Add data to map and key slice
			state := strings.Join(split_state, " ")
			data[state] = append(data[state], range_1, range_2)
			keys = append(keys, state)
		}
	}

	// -----------

	// Randomly select a state and a zipcode
	state := keys[rand.IntN(len(keys))]
	zipcode := rand.IntN(data[state][1]-data[state][0]) + data[state][0]

	ch <- state
	numch <- zipcode

	close(ch)
	close(numch)

	// -----------

}

// --------------------------------------------------------------------------------------------------------------------

func job_history(ch chan string, intch chan int, errch chan error) {
	// Generate a job history (based upon age)

	// -----------

	// Load data from file into memory
	const titles_file = "../data/job_titles.txt"

	job_titles, err := load_data(titles_file)
	if err != nil {
		errch <- fmt.Errorf("job_history %v", err)
	}

	// -----------

	// Select a random age and return it through a channel
	age := rand.IntN(85-18) + 18
	intch <- age

	first := true
	set := make(map[int]bool)
	start_date := 2024

	for age > 19 {

		// -----------

		// Select a unique job. Uses a set(map) to ensure unquieness
		job_title_index := rand.IntN(len(job_titles))
		func() {
			for {
				if !set[job_title_index] {
					set[job_title_index] = true
					break
				}
				job_title_index = rand.IntN(len(job_titles))
			}
		}()

		// -----------

		// Create a job title and time spend at that job
		random_choice := job_titles[job_title_index]
		time_spent := rand.IntN(age - 17)

		if first {
			ch <- fmt.Sprintf("%v - current\t%v", start_date-time_spent, random_choice)
			first = false
		} else {
			ch <- fmt.Sprintf("%v - %v\t%v", start_date-time_spent, start_date, random_choice)
		}

		start_date -= time_spent
		age -= time_spent

	}

	// -----------

	close(ch)
	close(intch)

}

// --------------------------------------------------------------------------------------------------------------------

func hobbies(total_hobbies int, ch chan string, errch chan error) {
	// Generate a random set of hobbies

	// -----------

	// Load data from text file into memory.
	const hobbies_file = "../data/hobbies.txt"

	hobby_data, err := load_data(hobbies_file)
	if err != nil {
		errch <- fmt.Errorf("hobbies: %v", err)
		return
	}

	// -----------

	// Return X random hobbies from hobby data.
	index_data := make(map[int]bool)

	// Use a set to pick X unique hobbies
	for len(index_data) <= total_hobbies {
		number := rand.IntN(len(hobby_data))
		if !index_data[number] {
			ch <- hobby_data[number]
			index_data[number] = true
		}
	}

	// -----------

	close(ch)
}

// --------------------------------------------------------------------------------------------------------------------

func load_data(location string) ([]string, error) {
	// Opens a file at a location

	// -----------

	// Open the file
	file, err := os.Open(location)
	if err != nil {
		return []string{}, fmt.Errorf("load_data: %v", err)
	}
	defer file.Close()

	// -----------

	// Create a scanner to read the lines of the text file
	var data []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	// -----------

	return data, nil
}

// --------------------------------------------------------------------------------------------------------------------
