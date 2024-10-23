package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Mouse struct {
	position rl.Vector2
	clicked  bool
}

// --------------------------------------------------------------------------------------------------------------------

func create_mouse(config *Config) Mouse {
	return Mouse{
		position: rl.Vector2{
			X: float32(config.width) / 2,
			Y: float32(config.height) / 2,
		},
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (m *Mouse) update_position(config *Config) {
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		m.position.X = float32(rl.GetMouseX())
		m.position.Y = float32(rl.GetMouseY())
		m.clicked = true
	} else {
		m.position.X = float32(rand.Intn(int(config.width)))
		m.position.Y = float32(rand.Intn(int(config.height)))
	}
}

// --------------------------------------------------------------------------------------------------------------------
