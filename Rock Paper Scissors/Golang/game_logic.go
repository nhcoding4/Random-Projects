// Logic for controlling the game state.

package main

import (
	"fmt"
	"math/rand"
)

// --------------------------------------------------------------------------------------------------------------------

// Total menu options
const total_options = 5

// --------------------------------------------------------------------------------------------------------------------

func check_winner(user_1_choice *Choice, user_2_choice *Choice,
	user_1_score *int, user_2_score *int, draws *int,
	player_1_name string, player_2_name string) {
	// Check for winner

	if user_1_choice.Check_Winner(*user_2_choice) {
		fmt.Printf("%v wins!\n", player_1_name)
		*user_1_score++
	} else if user_2_choice.Check_Winner(*user_1_choice) {
		fmt.Printf("%v wins!\n", player_2_name)
		*user_2_score++
	} else {
		fmt.Printf("\nDraw!\n")
		*draws++
	}
}

// --------------------------------------------------------------------------------------------------------------------

func end_game() bool {
	// End of game logic

	const choices = 2

	// -----------

	// Continue playing or exit
	fmt.Println("\nPress 1 to continue. Press 2 to exit...")
	continue_game := take_input(choices)

	// -----------

	// Clear the screen and move cursor back to the top left.
	reset_screen()

	// -----------

	// Exit the game if the user wanted to exit.
	return continue_game == 1
}

// --------------------------------------------------------------------------------------------------------------------

func human_vs_computer_menu() {
	// A human vs computer game.

	reset_screen()

	// Variables needed
	const human_players = 1
	const computer_name = "COMPUTER"

	var user_wins int
	var computer_wins int
	var draws int

	// -----------

	// Take user name
	name := get_name(human_players)

	// -----------

	for {

		// -----------

		// User and computer actions.
		user_action := take_action(name[0])
		computer_action := Choice(rand.Intn(total_options))
		fmt.Println("Computer picked:", computer_action)

		// -----------

		// Check for a winner.
		check_winner(&user_action, &computer_action, &user_wins, &computer_wins, &draws, name[0], computer_name)

		// -----------

		// End game actions
		if end_game() {
			break
		}

		// -----------

		// Display a scoreboard
		scoreboard(user_wins, computer_wins, draws, name[0], computer_name)
	}
}

// --------------------------------------------------------------------------------------------------------------------

func human_vs_human() {
	// A human vs human game.

	reset_screen()

	// -----------

	const max_players = 2

	// -----------

	// Get player names
	names := get_name(max_players)

	// Variables needed.
	var player_1_wins int
	var player_2_wins int
	var draws int

	for {

		// -----------

		// Take action from both players
		player_choices := []Choice{}
		for i := 1; i <= 2; i++ {
			fmt.Printf("%v, your turn!\n", names[i-1])
			player_choices = append(player_choices, take_action(names[i-1]))
			reset_screen()
		}

		// -----------

		// Print choices picked by users
		for i := 0; i < len(player_choices); i++ {
			fmt.Println(names[i], "picked", player_choices[i])
		}

		// -----------

		check_winner(&player_choices[0], &player_choices[1], &player_1_wins, &player_2_wins, &draws, names[0], names[1])

		// -----------

		// End game actions
		if end_game() {
			break
		}

		// -----------

		// Display a scoreboard
		scoreboard(player_1_wins, player_2_wins, draws, names[0], names[1])
	}
}

// --------------------------------------------------------------------------------------------------------------------

func scoreboard(player_score int, player_2_score int, draws int, player_1_name string, player_2_name string) {
	// Displays a scoreboard.

	total_games := player_score + player_2_score + draws

	fmt.Println("\nResults:\n-------")
	fmt.Println("Played:", total_games)
	fmt.Println(player_1_name, "wins", player_score)
	fmt.Println(player_2_name, "wins", player_2_score)
	fmt.Println("Draws:", draws)
	fmt.Println()
}

// --------------------------------------------------------------------------------------------------------------------
