const fs = require("fs")

class WordBank {
    // ----------------------------------------------------------------------------------------------------------------

    // File location.

    constructor(wordsFilePath) {
        this.location = wordsFilePath
    }

    // ----------------------------------------------------------------------------------------------------------------

    // Load words from the text file at the path provided to the word bank.

    get words() {
        try {
            const data = fs.readFileSync(this.location, "utf8")
            return data.split("\n")
        } catch (e) {
            console.log("Error loading words from file:", e.stack)
        }
    }

    // ----------------------------------------------------------------------------------------------------------------

    // Randomly select a word from the database.

    getWord() {

        return this.words[Math.floor(Math.random() * this.words.length)]
    }

    // ----------------------------------------------------------------------------------------------------------------

}

module.exports = WordBank
