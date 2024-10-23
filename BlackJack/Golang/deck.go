package main

import (
	"fmt"
	"math/rand/v2"
)

// --------------------------------------------------------------------------------------------------------------------

// Deck struct and methods.

// --------------------------------------------------------------------------------------------------------------------

type Deck struct {
	cards      []Card
	totalCards int
	totalDecks int
}

// ---------------------------------------------------------------------------------------------------------------------

func (d *Deck) remainingCards() int {
	return d.totalCards
}

// ---------------------------------------------------------------------------------------------------------------------

func (d *Deck) String() string {
	var cardsInDeck string
	for _, card := range d.cards {
		cardsInDeck = fmt.Sprintf("%v%v\n", cardsInDeck, card.String())
	}
	return cardsInDeck
}

// ---------------------------------------------------------------------------------------------------------------------

func (d *Deck) Populate() {
	cards := map[string]int{
		"Two":   2,
		"Three": 3,
		"Four":  4,
		"Five":  5,
		"Six":   6,
		"Seven": 7,
		"Eight": 8,
		"Nine":  9,
		"Ten":   10,
		"Jack":  10,
		"Queen": 10,
		"King":  10,
		"Ace":   11,
	}

	suits := [4]string{"Heart", "Club", "Spade", "Diamond"}

	for range d.totalDecks {
		for _, suit := range suits {
			for key, value := range cards {
				newCard := Card{value: value, name: key, suit: suit}
				d.cards = append(d.cards, newCard)
				d.totalCards++
			}
		}
	}
	d.shuffle()
}

// ---------------------------------------------------------------------------------------------------------------------

func (d *Deck) shuffle() {
	rand.Shuffle(d.totalCards, func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// ---------------------------------------------------------------------------------------------------------------------

func (d *Deck) drawCard() Card {
	const firstCard = 0
	const newFirst = 1

	if len(d.cards) == 0 {
		d.Populate()
	}
	card := d.cards[firstCard]
	d.cards = d.cards[newFirst:]
	d.totalCards--
	d.shuffle()
	return card
}

// ---------------------------------------------------------------------------------------------------------------------

func (d *Deck) resetDeck() {
	d.cards = []Card{}
	d.totalCards = 0
	d.Populate()
}

// ---------------------------------------------------------------------------------------------------------------------

// Deck helper functions.

// ---------------------------------------------------------------------------------------------------------------------

func createDeck(totalDecks int) *Deck {
	newDeck := &Deck{totalDecks: totalDecks}
	newDeck.Populate()
	return newDeck
}

// --------------------------------------------------------------------------------------------------------------------
