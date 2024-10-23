package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// --------------------------------------------------------------------------------------------------------------------

func main() {
	// Object creatio

	var config = Config{
		width:      1000,
		height:     1000,
		target_fps: 144,
		title:      "Explosions",
	}
	config.init_window()

	var manager = Explosion_Manager{}
	manager.init(&config)

	for !rl.WindowShouldClose() {

		// Updates
		config.resize_window()
		manager.fire_projectile(&config)
		manager.update(&config)

		// Drawing
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		manager.draw()

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

// --------------------------------------------------------------------------------------------------------------------
