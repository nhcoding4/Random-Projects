// Counts the character frequency found inside a textfile.

// --------------------------------------------------------------------------------------------------------------------

const fs = require("fs")

// --------------------------------------------------------------------------------------------------------------------

function main() {

    const dataLocation = "frankenstien.txt"

    const fileData = loadData(dataLocation)

    const words = splitData(fileData)

    const characterCount = countCharacters(words)

    const reportString = report(characterCount)
    console.log(reportString)
}

// --------------------------------------------------------------------------------------------------------------------

// Creates a string reporting the data found. 

function report(data) {
    const sortedData = Object.fromEntries(
        Object.entries(data).sort(([, a], [, b]) => a - b)
    )

    let reportString = ""
    for (const key in sortedData) {
        reportString += `The character '${key}' was found: ${sortedData[key]} times.\n`
    }
    return reportString
}

// --------------------------------------------------------------------------------------------------------------------

// Sum character total in data.

function countCharacters(data) {
    let characters = {}

    for (const word of data) {
        for (let letter of word) {
            if (letter.toLowerCase() == letter.toUpperCase()) {
                continue
            }
            letter = letter.toLowerCase()
            if (!characters[letter]) {
                characters[letter] = 1
            } else {
                characters[letter]++
            }
        }
    }
    return characters
}

// --------------------------------------------------------------------------------------------------------------------

// Load the textfile into memory

function loadData(fileLocation) {

    try {
        const data = fs.readFileSync(fileLocation, "utf8")
        return data
    } catch (e) {
        console.log("Error", e.stack)
    }
    return ""
}

// --------------------------------------------------------------------------------------------------------------------

// Split data using regex

function splitData(data) {
    const seperators = [" ", "\n"]
    const words = data.split(new RegExp(seperators.join("|"), "g"))
    return words
}

// --------------------------------------------------------------------------------------------------------------------

main()
