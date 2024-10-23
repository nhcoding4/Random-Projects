package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Effect struct {
	width               int32
	height              int32
	linkDistance        int32
	particles           []*Particle
	mouse               *Mouse
	totalParticles      int32
	powerPushMultiplier float64
}

// ---------------------------------------------------------------------------------------------------------------------

func (e *Effect) init() {
	e.createParticles()
}

// ---------------------------------------------------------------------------------------------------------------------

func (e *Effect) calculateDistance(particle1, particle2 *Particle) float64 {
	xDistance := particle1.x - particle2.x
	yDistance := particle1.y - particle2.y
	return math.Hypot(float64(xDistance), float64(yDistance))
}

// ---------------------------------------------------------------------------------------------------------------------

func (e *Effect) connectParticles() {
	length := len(e.particles)

	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if j == i {
				continue
			}
			distance := e.calculateDistance(e.particles[i], e.particles[j])

			if distance < float64(e.linkDistance) {
				e.drawLines(&distance, e.particles[i], e.particles[j])
			}
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (e *Effect) createParticles() {
	for range e.totalParticles {
		newParticle := Particle{
			width:               e.width,
			height:              e.height,
			mouse:               e.mouse,
			powerPushMultiplier: e.powerPushMultiplier,
		}
		newParticle.init()
		e.particles = append(e.particles, &newParticle)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (e *Effect) draw() {
	e.connectParticles()
	for i := 0; i < len(e.particles); i++ {
		e.particles[i].draw()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (e *Effect) drawLines(distance *float64, particle1, particle2 *Particle) {
	opacity := 1 - (*distance / float64(e.linkDistance))

	rl.DrawLineEx(
		rl.Vector2{X: float32(particle1.x), Y: float32(particle1.y)},
		rl.Vector2{X: float32(particle2.x), Y: float32(particle2.y)},
		float32(2.5),
		rl.Color{R: 255, G: 255, B: 255, A: uint8(255 * opacity)},
	)
}

// ---------------------------------------------------------------------------------------------------------------------

func (e *Effect) setPosition() {
	for i := 0; i < len(e.particles); i++ {
		e.particles[i].setPosition()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (e *Effect) update() {
	for i := 0; i < len(e.particles); i++ {
		e.particles[i].update()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (e *Effect) updateWindowSize() {
	e.width = int32(rl.GetScreenWidth())
	e.height = int32(rl.GetScreenHeight())

	for i := 0; i < len(e.particles); i++ {
		e.particles[i].updateWindowSize()
	}
}

// ---------------------------------------------------------------------------------------------------------------------
