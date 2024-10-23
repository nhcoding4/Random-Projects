package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
	"path/filepath"
)

// ---------------------------------------------------------------------------------------------------------------------

// Helpers functions

// ---------------------------------------------------------------------------------------------------------------------

// Creates and returns a pointer to a game struct.

func gameSetup() *game {
	gameStruct := &game{
		title:            "Snake",
		targetFPS:        144,
		cellSize:         30,
		cellCount:        25,
		lastUpdate:       0,
		updateInterval:   0.2,
		borderOffset:     75,
		backgroundColour: rl.Black,
	}
	unpackFiles()
	gameStruct.init()
	return gameStruct
}

// --------------------------------------------------------------------------------------------------------------------

// Returns the path of the score file.

func filePaths() (string, string, error) {
	const directoryName = "highScore"
	const fileName = "highScore.txt"

	rootDirectory, err := os.Getwd()
	if err != nil {
		return "", "", err
	}

	directoryLocation := filepath.Join(rootDirectory, directoryName)
	filePath := filepath.Join(directoryLocation, fileName)

	return directoryLocation, filePath, nil
}

// --------------------------------------------------------------------------------------------------------------------

// Unpacks the embedded files.

func unpackFiles() {

	wallRaw, err := resources.ReadFile("resources/wall.mp3")
	if err != nil {
		fmt.Println(err)
	}
	eatRaw, err := resources.ReadFile("resources/eat.mp3")
	if err != nil {
		fmt.Println(err)
	}
	foodRaw, err := resources.ReadFile("resources/food.png")
	if err != nil {
		fmt.Println(err)
	}

	err = os.Mkdir("temp", os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile("temp/eat.mp3", eatRaw, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("temp/wall.mp3", wallRaw, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("temp/food.png", foodRaw, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

// ---------------------------------------------------------------------------------------------------------------------
