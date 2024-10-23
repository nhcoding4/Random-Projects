package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"os"
)

// ---------------------------------------------------------------------------------------------------------------------

// Struct and methods relating to game state.

// ---------------------------------------------------------------------------------------------------------------------

type Game struct {
	players []*Player
	dealer  *Dealer
	deck    *Deck
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) setup() {
	const totalDecks = 8
	g.deck = createDeck(totalDecks)
	clearScreen()
	g.createPlayers()
	g.dealer = createDealer(g.deck)
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) createPlayers() {
	const prompt = "How many players would you like to add?"
	const minimum = 1
	const maximum = 4
	totalPlayers := takeInputInt(prompt, minimum, maximum)

	for i := 0; i < totalPlayers; i++ {
		playerNumber := i + 1
		g.addPlayer(playerNumber)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) addPlayer(playerNumber int) {
	const minChips = 1
	const maxChips = 100_000

	playerPrompt := fmt.Sprintf("Enter player %v's name", playerNumber)
	playerName := takeInputString(playerPrompt)

	chipsPrompt := fmt.Sprintf("Enter the chips player %v is buying in for", playerName)
	playerChips := takeInputFloat(chipsPrompt, minChips, maxChips)

	g.players = append(g.players, createNewPlayer(playerName, playerChips, g.deck))
	printSpacer()
}

// --------------------------------------------------------------------------------------------------------------------

func (g *Game) bettingRound() {
	for _, player := range g.players {
		if player.chips > 0 {
			player.makeBet()
			printSpacer()
		} else {
			fmt.Printf("%v has no chips and cannot make a bet.\n", player.playerName)
			printSpacer()
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) drawCards() {
	const hit = 1
	const stand = 2
	const cards = 1
	const prompt = "Press 1 to Draw. Press 2 to Stand:"

	for _, player := range g.players {
		g.dealer.initialCardDisplay()
		player.printCards()
		for {
			if player.stopTurn() {
				break
			}
			if takeInputInt(prompt, hit, stand) == hit {
				player.getPlayerCard(cards, g.deck)
				player.printCards()
			} else {
				break
			}
		}
		printSpacer()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) dealerActions() {
	const card = 1
	const dealerMustHit = 16

	g.dealer.printCards()

	for g.dealer.score <= dealerMustHit {
		g.dealer.getPlayerCard(card, g.deck)
		g.dealer.printCards()
		if g.dealer.stopTurn() {
			break
		}
	}
	printSpacer()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) distributeChips() {
	for _, player := range g.players {
		winAmount := player.bet * 2.0
		blackJackAmount := player.bet * 2.25

		if !player.bust {
			if checkDraw(g.dealer.blackJack, player.blackJack, player.score, g.dealer.score) {
				player.chips += player.bet
				fmt.Printf("%v push | Total Chips: %v\n", player.playerName, player.chips)
				printSpacer()

			} else if checkWin(player.blackJack, player.score, g.dealer.score) {
				player.chips += winAmount
				fmt.Printf("%v wins %v chips | Total Chips: %v\n", player.playerName, winAmount, player.chips)
				printSpacer()

			} else if player.blackJack {
				player.chips += blackJackAmount
				fmt.Printf("%v has BLACKJACK and wins %v chips | Total Chips: %v\n",
					player.playerName, blackJackAmount, player.chips)
				printSpacer()
			}
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) postGame() bool {
	if !continueGame() {
		return false
	}
	g.reBuyRemovePlayers()
	g.resetPlayers()
	g.deck.resetDeck()

	return true
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) removePlayer(index int) {
	if len(g.players)-1 == 0 {
		fmt.Println("There is no one else left to play. Game over!")
		os.Exit(1)
	}

	newPlayerSlice := make([]*Player, len(g.players)-2)
	for i := range g.players {
		if i == index {
			continue
		}
		newPlayerSlice = append(newPlayerSlice, g.players[i])
	}
	g.players = newPlayerSlice
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) resetPlayers() {
	g.dealer.postGameActions()

	for _, player := range g.players {
		player.postGameActions()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) reBuyRemovePlayers() {
	for i, player := range g.players {
		if player.chips == 0 {
			keepPlayer := player.reBuy()
			if !keepPlayer {
				g.removePlayer(i)
			}
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) dealNewCards() {
	const newCards = 2

	g.dealer.getPlayerCard(newCards, g.deck)

	for _, player := range g.players {
		player.getPlayerCard(newCards, g.deck)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) gameLoop() {
	g.setup()

	for {
		clearScreen()
		g.bettingRound()
		g.drawCards()
		g.dealerActions()
		g.distributeChips()
		if !g.postGame() {
			break
		}
		g.dealNewCards()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Helper functions related to game structs methods.

// ----------------------------------------------------------------------------------------------------------------------

func checkWin(playerBlackJack bool, playerScore, dealerScore int) bool {
	return !playerBlackJack && playerScore > dealerScore
}

// ---------------------------------------------------------------------------------------------------------------------

func checkDraw(dealerBlackJack, playerBlackJack bool, playerScore, dealerScore int) bool {
	return dealerBlackJack && playerBlackJack || playerScore == dealerScore
}

// ---------------------------------------------------------------------------------------------------------------------

func continueGame() bool {
	const endGamePrompt = "Would you like to continue playing? 1 to continue, 2 to quit."
	const continueGame = 1
	const exitGame = 2

	if takeInputInt(endGamePrompt, continueGame, exitGame) == exitGame {
		return false
	}
	return true
}

// --------------------------------------------------------------------------------------------------------------------

func clearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

// --------------------------------------------------------------------------------------------------------------------
