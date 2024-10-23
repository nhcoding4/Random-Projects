// A higher/lower game with X number of players.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

// --------------------------------------------------------------------------------------------------------------------

type player struct {
	name   string
	number int
}

func (p *player) select_number() {
	p.number = rand.Intn(100)
}

// --------------------------------------------------------------------------------------------------------------------

func main() {

	// Take input
	total_players := func() int {
		for {
			fmt.Println("Enter the amount of players to battle it out:")

			// Create a scanner to read input.
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			err := scanner.Err()
			if err != nil {
				log.Fatal(err)
			}

			// Attempt to convert input into a valid integer.
			number, err := strconv.Atoi(scanner.Text())
			if err != nil || number <= 0 {
				fmt.Println("Invalid input. Please enter a positive whole number.")
			} else {
				return number
			}
		}
	}()

	// -----------

	// Create the amount of players to battle
	players := []player{}
	for i := 0; i < total_players; i++ {
		players = append(players, player{name: fmt.Sprintf("%v", i+1)})
	}

	// -----------

	// Find a winner
	current_round := 1
	for len(players) > 1 {
		fmt.Println("Round", current_round, "\n-------")
		fmt.Println()
		players = start_round(players)
		current_round++
	}

	fmt.Println(players[0].name, "Has won!")
}

// --------------------------------------------------------------------------------------------------------------------

func start_round(players_remaining []player) []player {
	// Round logic.

	// -----------

	// Variables needed for each round.
	var waitgroup sync.WaitGroup
	rounds := int(len(players_remaining) / 2)
	i, j := 0, 1
	winners := make(chan player, rounds)

	// -----------

	// Use a go routine to get winners
	for rounds > 0 {
		waitgroup.Add(1)
		go get_winner(&waitgroup, players_remaining[i], players_remaining[j], winners)
		i += 2
		j += 2
		rounds--
	}

	waitgroup.Wait()

	// -----------

	// Get victors of the round
	victors := []player{}
	for range len(winners) {
		round_winner := <-winners
		victors = append(victors, round_winner)
	}

	// Add leftover players
	for i < len(players_remaining) {
		victors = append(victors, players_remaining[i])
		i++
	}

	// -----------

	close(winners)
	return victors
}

// --------------------------------------------------------------------------------------------------------------------

func get_winner(wg *sync.WaitGroup, player_1 player, player_2 player, ch chan player) {
	// Takes 2 players and picks a winner

	// -----------

	// Tell the waitgroup the current worker has finished on function return.
	defer wg.Done()

	// ----------

	// Generate random numbers
	for {
		player_1.select_number()
		player_2.select_number()

		if player_1.number != player_2.number {
			break
		}
	}

	// ----------

	// Select a winner ()
	if player_1.number > player_2.number {
		fmt.Printf("Player: %v with %v VS Player: %v with %v\n%v wins!\n\n",
			player_1.name, player_1.number, player_2.name, player_2.number, player_1.name,
		)
		ch <- player_1
		return
	}

	fmt.Printf("Player: %v with %v VS Player: %v with %v\n%v wins!\n\n",
		player_1.name, player_1.number, player_2.name, player_2.number, player_2.name,
	)

	ch <- player_2
}

// --------------------------------------------------------------------------------------------------------------------
