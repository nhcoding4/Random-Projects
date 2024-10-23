package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Mouse struct {
	activated bool
	radius    int32
	x         int32
	y         int32
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *Mouse) activationStatus() *bool {
	return &m.activated
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

func (m *Mouse) getLocation() (*int32, *int32) {
	return &m.x, &m.y
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *Mouse) getRadius() *int32 {
	return &m.radius
}

// ---------------------------------------------------------------------------------------------------------------------

func (m *Mouse) update() {
	m.x = rl.GetMouseX()
	m.y = rl.GetMouseY()
}

// ---------------------------------------------------------------------------------------------------------------------
