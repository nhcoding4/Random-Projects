package main

import "fmt"

// ---------------------------------------------------------------------------------------------------------------------

// Dealer struct and methods.

// ---------------------------------------------------------------------------------------------------------------------

type Dealer struct {
	*Player
}

// ---------------------------------------------------------------------------------------------------------------------

func (d *Dealer) initialCardDisplay() {
	fmt.Printf("The dealer is showing %v | Score: %v\n", d.playerCards[0].String(), d.playerCards[0].value)
}

// ---------------------------------------------------------------------------------------------------------------------

func createDealer(deck *Deck) *Dealer {
	const name = "Dealer"
	const chips = 1

	newDealer := &Dealer{}
	newDealer.Player = createNewPlayer(name, chips, deck)
	return newDealer
}

// ---------------------------------------------------------------------------------------------------------------------
