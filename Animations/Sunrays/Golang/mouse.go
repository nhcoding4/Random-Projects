package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Mouse struct {
	activated bool
	radius    int32
	x         int32
	y         int32
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *Mouse) checkLeftClickDown() {
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		m.activated = true
	} else {
		m.activated = false
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *Mouse) update() {
	m.x = rl.GetMouseX()
	m.y = rl.GetMouseY()
}

// ---------------------------------------------------------------------------------------------------------------------
