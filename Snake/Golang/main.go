package main

import (
	"embed"
	_ "embed"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// --------------------------------------------------------------------------------------------------------------------

//go:embed resources/*
var resources embed.FS

// --------------------------------------------------------------------------------------------------------------------

func main() {

	game := gameSetup()

	// ----------------------------------------------------

	// Main game loop.

	for rl.WindowShouldClose() == false {
		rl.BeginDrawing()
		game.update()

		rl.ClearBackground(game.backgroundColour)
		game.draw()

		rl.EndDrawing()
	}

	// ----------------------------------------------------

	rl.CloseAudioDevice()
	rl.CloseWindow()
	game.exitActions()
}

// ---------------------------------------------------------------------------------------------------------------------
