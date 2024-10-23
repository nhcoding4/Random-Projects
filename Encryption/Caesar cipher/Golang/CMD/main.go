package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"

	"github.com/inancgumus/screen"
)

// --------------------------------------------------------------------------------------------------------------------

func main() {

	const encrypt = 1
	const decrypt = 2

	// -----------

	for {

		// Display options
		screen.Clear()
		screen.MoveTopLeft()
		fmt.Println("---> CAESAR'S CIPHER <---")
		fmt.Println("1) Encrypt message")
		fmt.Println("2) Decrypt message")
		fmt.Println("3) Exit program")

		// -----------

		// Take user input. Check for a valid selection.
		prompt := "Select an option"
		error_message := "Please enter an integer between 1 and 3"
		choice := func() int {
			for {
				user_input := take_input_integer(prompt, error_message)
				if user_input > 3 {
					fmt.Println(error_message)
				} else {
					return user_input
				}
			}
		}()

		// -----------

		// Execute choice.
		switch choice {

		case 1:
			// Attempt to encrypt a message.
			message := "Please enter a message to encrypt..."
			option_wrapper(message, encrypt)

			// -----------

		case 2:
			// Attempt to decrypt a message.
			message := "Please enter a message to decrypt..."
			option_wrapper(message, decrypt)

			// -----------

		case 3:
			// Exit the program.
			screen.Clear()
			screen.MoveTopLeft()
			os.Exit(1)

			// -----------
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------

func take_input(prompt string) (string, error) {
	// Takes a string input from the user.

	// Promt the user for input.
	fmt.Println(prompt)

	// -----------

	// Create a new scanner to read input from the user.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// -----------

	// Check for errors.
	err := scanner.Err()
	if err != nil {
		return "", err
	}

	// -----------

	return scanner.Text(), nil
}

// --------------------------------------------------------------------------------------------------------------------

func take_input_integer(prompt string, error_message string) int {
	// Take an interger from the user.

	for {

		// Take input from the user.
		user_input, err := take_input(prompt)
		if err != nil {
			log.Fatal(err)
		}

		// -----------

		// Attempt to convert it to a number.
		number, err := strconv.Atoi(user_input)
		if err != nil || number <= 0 {
			fmt.Println(error_message)
		} else {
			return number
		}

		// -----------
	}
}

// --------------------------------------------------------------------------------------------------------------------

func encrypt_decrypt(message string, option int) string {
	// Encrypts/Decrypts a string using Caesar's Cipher.

	// Take an offset (needed for encryption) from the user.
	prompt := "Enter an offset"
	error_message := "Please enter a positive integer"
	offset := take_input_integer(prompt, error_message)

	// -----------

	var adjusted_message string

	// -----------

	// Encrypt the message.
	for _, char := range message {

		// Add in non A-Z/a-z letters as they are.
		if !unicode.IsLetter(char) {
			adjusted_message += string(char)
			continue
		}

		// -----------

		// Upper and lower limits of conversion.
		lower_bound := 'a'
		upper_bound := 'z'

		// Deal with uppercase letters
		if char >= 'A' && char <= 'Z' {
			lower_bound = 'A'
			upper_bound = 'Z'
		}

		// -----------

		switch option {

		case 1:
			// If the character is moved beyond the alphabet, move the character to 'a'.
			for i := 0; i < offset; i++ {
				char++
				if char > upper_bound {
					char = lower_bound
				}
			}

			// -----------

		case 2:
			// if the character is moved beyond the alphabet, move the character to 'z'.
			for i := 0; i < offset; i++ {
				char--
				if char < lower_bound {
					char = upper_bound
				}
			}

			// -----------
		}

		// Add the rune to the string.
		adjusted_message += string(char)
	}

	return adjusted_message
}

// --------------------------------------------------------------------------------------------------------------------

func continue_program() {
	// Pauses the program until the user presses enter.

	fmt.Println("Press 'Enter' to continue...")
	var message string
	fmt.Scanf("%v", &message)
}

// --------------------------------------------------------------------------------------------------------------------

func option_wrapper(prompt string, option int) {

	screen.Clear()
	screen.MoveTopLeft()
	message, err := take_input(prompt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(encrypt_decrypt(message, option))
	continue_program()
}
