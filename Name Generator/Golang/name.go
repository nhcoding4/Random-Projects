// Main function.

package main

import (
	"fmt"
)

// --------------------------------------------------------------------------------------------------------------------

type Person_CV struct {
	name            string
	sex             string
	age             int
	social_Security string
	location        string
	zipcode         int
	hobbies         []string
	job_history     []string
}

func (person *Person_CV) Print() {
	// Prints the contents of the struct to the console.

	fmt.Println("\nPersonal Data\n------------")
	fmt.Println("Name:", person.name)
	fmt.Println("Sex:\t", person.sex)
	fmt.Println("Age:\t", person.age)
	fmt.Println("Social Security:", person.social_Security)

	fmt.Println("\nLocation:\n--------")
	fmt.Println("City:\t", person.location)
	fmt.Println("Zipcode:", person.zipcode)

	fmt.Println("\nEmployment History\n------------------")
	for _, value := range person.job_history {
		fmt.Println(value)
	}

	fmt.Println("\nInterests\n---------")
	for i, value := range person.hobbies {
		fmt.Println(i+1, "\t", value)
	}
	fmt.Println()
}

// --------------------------------------------------------------------------------------------------------------------

func main() {

	// Randomly generate credentials of a person and print it to the screen.
	created_person := collect_data()
	created_person.Print()
}

// --------------------------------------------------------------------------------------------------------------------
