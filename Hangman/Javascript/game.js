const prompt = require("prompt-sync")({ sigint: true })
const WordBank = require("./wordbank.js")
const art = require("./art.js")

class Game {
    // ----------------------------------------------------------------------------------------------------------------

    // Game variables

    constructor(wordsFilePath) {
        this.wordbank = new WordBank(wordsFilePath)
    }
    lives = 7
    word = ""
    underscores = ""
    guess = ""
    history = []
    incorrect = []

    // ----------------------------------------------------------------------------------------------------------------

    // Main loop.

    mainLoop() {
        while (true) {
            this.roundStart()
            if (this.checkFatal()) {
                return
            }

            while (this.lives > 0) {
                console.clear()
                this.displayArt()
                this.displayInformation()
                this.guess = prompt("Make a guess: ")

                if (!this.checkCharacters()) {
                    this.reduceLives()
                    if (this.lives == 0) {
                        console.log(`\nGame over! The word was: ${this.word}`)
                        break
                    }
                }
                if (this.checkWin()) {
                    console.log(`\n${this.word}\nYou got the word! You win!`)
                    break
                }
                console.log()
            }

            console.log()
            if (!this.roundEnd()) {
                console.log("Thanks for playing!")
                break
            }
        }
    }

    // ----------------------------------------------------------------------------------------------------------------

    // Checks the word for characters against the first character the user entered. Reveals correct guesses.

    checkCharacters() {
        let found = false
        let newString = ""
        const userGuess = this.guess[0].toLowerCase()
        this.history.push(userGuess)

        for (let i = 0; i < this.word.length; i++) {
            if (userGuess == this.word[i]) {
                newString += userGuess
                found = true
            } else {
                if (this.underscores[i] != "_") {
                    newString += this.underscores[i]
                } else {
                    newString += "_"
                }
            }
        }

        this.underscores = newString
        return found
    }

    // ----------------------------------------------------------------------------------------------------------------

    // Checks if there is a valid word bank.

    checkFatal() {
        if (this.wordbank.length == 0) {
            console.log("There was an error loading the words from the words.txt file.")
            return true
        }
        return false
    }

    // ----------------------------------------------------------------------------------------------------------------

    // Check if user has guessed the word.

    checkWin() {
        if (this.underscores == this.word) {
            return true
        }
        return false
    }

    // ----------------------------------------------------------------------------------------------------------------

    // Convert the selected word into underscores.

    createUnderscores() {
        this.underscores = ""

        for (let i = 0; i < this.word.length; i++) {
            this.underscores += "_"
        }
    }

    // ----------------------------------------------------------------------------------------------------------------

    // Game art. Changes depending on the amount of lives remaining.

    displayArt() {
        if (this.lives != 7) {
            console.log(art(this.lives))
        }
    }

    // ----------------------------------------------------------------------------------------------------------------


    // Displays game state information to the user.

    displayInformation() {
        console.log(this.underscores)
        console.log(`You have ${this.lives} lives remaining.`)
        console.log(`History: ${this.history} | Incorrect: ${this.incorrect}`)
    }

    // ----------------------------------------------------------------------------------------------------------------

    // Removes a life if the user makes a unique wrong guess.

    reduceLives() {
        let found = false
        const userGuess = this.guess[0].toLowerCase()

        for (let letter of this.incorrect) {
            if (letter == userGuess) {
                found = true
            }
        }
        if (!found) {
            this.lives--
            this.incorrect.push(userGuess)
        }
    }

    // ----------------------------------------------------------------------------------------------------------------

    // Continue/stop game.

    roundEnd() {
        let choice = prompt("Would you like to play another round? (yes/no): ").toLowerCase()

        if (choice == "yes" || choice == "y") {
            return true
        }
        return false
    }

    // ----------------------------------------------------------------------------------------------------------------

    // Set variables to starting state.

    roundStart() {
        this.lives = 7
        this.word = this.wordbank.getWord()
        this.createUnderscores()
        this.history = []
        this.incorrect = []
    }

    // ----------------------------------------------------------------------------------------------------------------

}

module.exports = Game
