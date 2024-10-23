/*
Single threaded Benchmarks:
	2px size = 8fps
	5px size = 55fps
	7px size = 105fps
	10px size = 200fps
*/

package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// --------------------------------------------------------------------------------------------------------------------

func main() {

	// -------------------------------

	screen_width := 1920
	screen_height := 1080
	cells_size := 2
	const title = "Game of Life"
	target_fps := int32(1000)
	rl.InitWindow(int32(screen_width), int32(screen_height), title)
	rl.SetTargetFPS(target_fps)

	// Create a simulation object. This controls the game.
	simulation := Simulation{cell_size: cells_size, height: screen_height, width: screen_width, run: false}
	simulation.Populate_Grid()

	// -------------------------------

	// Mainloop.

	for !rl.WindowShouldClose() {

		// -------------------------------
		// User Input.

		key := rl.GetKeyPressed()
		simulation.Input_Actions(key, &target_fps)

		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			simulation.Mouse_Function()
		}

		// -------------------------------

		simulation.Update()

		// -------------------------------

		rl.BeginDrawing()

		rl.ClearBackground(rl.DarkGray)

		simulation.grid.Draw()
		simulation.fps()

		rl.EndDrawing()

		// -------------------------------
	}

	rl.CloseWindow()
}

// --------------------------------------------------------------------------------------------------------------------
