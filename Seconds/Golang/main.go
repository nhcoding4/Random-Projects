package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"atomicgo.dev/cursor"
)

// --------------------------------------------------------------------------------------------------------------------

func main() {

	// Get the target year.
	year := day_year("Enter a year:")

	// Get the target month.
	month := month()

	// Get the target day. Check for invalid days. If an invalid day has been entered then keep prompting.
	day := func() int {

		for {

			// Take a date from the user.
			entered_day := day_year("Enter a day:")

			// ----------

			// 0 or negative day values is invalid.
			if entered_day < 1 {
				fmt.Println("Invalid day entered")
				continue
			}

			// ----------

			// Check for valid day depending on the month/year entered.

			var invalid_date bool
			switch month {

			// Months with 31 days.
			case 1, 3, 5, 7, 8, 10, 12:

				if entered_day > 31 {
					invalid_date = true
				}

				// Months with 30 days.
			case 4, 6, 9, 11:
				if entered_day > 30 {
					invalid_date = true
				}

				// Leap years and February.
			case 2:
				if (year%400 == 0) || ((year%100 != 0) && (year%4 == 0)) {
					if entered_day > 29 {
						invalid_date = true
					}
				} else {
					if entered_day > 28 {
						invalid_date = true
					}
				}
			}

			// ----------

			// Reprompt on garbage value
			if invalid_date {
				fmt.Println("Invalid date entered")
				continue
			}

			// ----------

			return entered_day
		}
	}()

	// ----------

	// Display information to user on valid input.

	go convert(year, month, day)
	_ = take_input("")

}

// --------------------------------------------------------------------------------------------------------------------

func convert(year int, month int, day int) {
	// Displays the time (in seconds) since/to the entered date. Updates in real time.

	// ----------

	fmt.Println("There has been:")

	for {

		// ----------

		// Get the time when the program is ran and create a time object from the user input.
		now := time.Now()
		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

		// ----------

		// Calculate the difference in time.
		difference := now.Sub(date)

		// ----------

		// Display and update the time displayed.
		var connector string

		if difference < 0 {
			difference *= -1
			connector = " seconds until "
		} else {
			connector = " seconds since "
		}

		fmt.Print(math.Round(difference.Seconds()), connector, date.Day(), date.Month(), date.Year())
		fmt.Println("\nPress 'ENTER' to Exit.")
		time.Sleep(time.Second)
		cursor.UpAndClear(2)
	}
}

// --------------------------------------------------------------------------------------------------------------------

func take_input(prompt string) string {
	// Displays a prompt to the user and takes input. Keeps prompting until a valid integer has been entered.

	// ----------

	fmt.Println(prompt)

	// ----------

	// Create a new scanner and return input from stdin.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()

}

// --------------------------------------------------------------------------------------------------------------------

func month() int {
	// Allows the user to enter a month using names.

	// ----------

	// Months of the year.
	months := [12]string{
		"January",
		"Feburary",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	// ----------

	// Get a month from the user.
	for {

		user_input := take_input("Enter a month:")

		// User if no error, user has entered a numerical month.
		value, err := strconv.Atoi(user_input)

		// ----------

		// See if the use has entered a month by name.
		if err != nil {

			converted_value, err := func(month_name string) (int, error) {

				// Convert user input to lower-case to prevent parsing errors.
				month_name = strings.ToLower(month_name)

				// Check user input against each month for a valid input.
				for i, month := range months {
					if month_name == strings.ToLower(month) {
						return i + 1, nil
					}
				}

				// Garbage value entered.
				return -1, errors.New("month: cannot convert input to a month")

			}(user_input)

			// ----------

			// Keep prompting use if they have entered some garbage value.
			if err != nil {
				fmt.Println(err)
				continue
			}

			return converted_value
		}

		// ----------

		// Check for garbage value
		if value < 1 || value > 12 {
			fmt.Println("Invalid input. Please enter a valid month or a number between 1 and 12.")
			continue
		}

		// ----------

		return value
	}
}

// --------------------------------------------------------------------------------------------------------------------

func day_year(prompt string) int {
	// Used for taking day and year values from the user.

	// ----------

	// Keep prompting user for input until they enter a valid integer.
	for {

		// Get user input.
		user_input := take_input(prompt)

		// Attempt to conver the input into an integer.
		integer_value, err := strconv.Atoi(user_input)

		// Check for error.
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}

		// ----------

		return integer_value
	}
}
