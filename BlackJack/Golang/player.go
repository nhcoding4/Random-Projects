package main

import "fmt"

// ---------------------------------------------------------------------------------------------------------------------

// Player struct and methods.

// ---------------------------------------------------------------------------------------------------------------------

type Player struct {
	playerName  string
	playerCards []Card
	chips       float64
	bet         float64
	score       int
	bust        bool
	blackJack   bool
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Player) getPlayerCard(amount int, currentDeck *Deck) {
	for i := 0; i < amount; i++ {
		newCard := currentDeck.drawCard()
		p.playerCards = append(p.playerCards, newCard)
		p.score += newCard.value
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Player) printCards() {
	cardString := fmt.Sprintf("%v", p.playerName)

	for _, card := range p.playerCards {
		cardString = fmt.Sprintf("%v | %v", cardString, card.String())
	}
	fmt.Printf("%v | Score: %v\n", cardString, p.score)
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Player) makeBet() {
	const minimum = 1

	p.bet = 0

	if p.chips > 0 {
		prompt := fmt.Sprintf("%v, you have %v chips. Place your bet:", p.playerName, p.chips)
		betAmount := takeInputFloat(prompt, minimum, p.chips)
		p.chips -= betAmount
		p.bet = betAmount
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Player) checkAce() {
	for _, card := range p.playerCards {
		card.convertAceLow()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Player) reBuy() bool {
	reBuyPrompt := fmt.Sprintf("%v would you like to rebuy? Press 1 for yes, 2 to quit.", p.playerName)
	const yes = 1
	const no = 2
	const chipsPrompt = "How much would you like to addon?"
	const minChips = 1
	const maxChips = 1000

	if takeInputInt(reBuyPrompt, yes, no) == yes {
		p.chips += takeInputFloat(chipsPrompt, minChips, maxChips)
		return true
	}
	return false
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Player) stopTurn() bool {
	if p.stopChecks() {
		if !p.blackJack {
			p.checkAce()
			if p.stopChecks() {
				fmt.Printf("%v has bust!\n", p.playerName)
				return true
			}
			fmt.Printf("%v has blackJack\n!", p.playerName)
			return false
		}
		return true
	}
	return false
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Player) stopChecks() bool {
	const maxScore = 21
	const blackJackCards = 2

	if p.score == maxScore {
		if len(p.playerCards) == blackJackCards {
			p.blackJack = true
		}
		return true
	}
	if p.score > maxScore {
		p.bust = true
		return true
	}
	return false
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Player) postGameActions() {
	const resetValue = 0

	p.playerCards = []Card{}
	p.blackJack = false
	p.bust = false
	p.bet = resetValue
	p.score = resetValue
}

// ---------------------------------------------------------------------------------------------------------------------

// Helper functions related to player structs or methods.

// ---------------------------------------------------------------------------------------------------------------------

func createNewPlayer(name string, chips float64, deck *Deck) *Player {
	const startingCards = 2

	newPlayer := &Player{playerName: name}
	newPlayer.chips = chips
	newPlayer.getPlayerCard(startingCards, deck)
	return newPlayer
}

// ---------------------------------------------------------------------------------------------------------------------
