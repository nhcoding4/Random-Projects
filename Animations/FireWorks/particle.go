package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Particle struct {
	width, height      int32
	velocity, position rl.Vector2
	color              rl.Color
	active             bool
}

// --------------------------------------------------------------------------------------------------------------------

func (p *Particle) init(position *rl.Vector2) {
	p.width = int32(rand.Intn(5) + 5)
	p.height = int32(rand.Intn(5) + 5)
	p.position = *position
	p.velocity = rl.Vector2{
		X: float32(rand.Intn(6)-3) + rand.Float32(),
		Y: float32(rand.Intn(6)-3) + rand.Float32(),
	}
	p.color = rl.Color{
		R: uint8(rand.Intn(255)),
		G: uint8(rand.Intn(255)),
		B: uint8(rand.Intn(255)),
		A: 255,
	}
	p.active = true
}

// --------------------------------------------------------------------------------------------------------------------

func create_particle(position *rl.Vector2) Particle {
	var new_particle = Particle{}
	new_particle.init(position)
	return new_particle
}

// --------------------------------------------------------------------------------------------------------------------

func (p *Particle) draw() {
	if p.active {
		rl.DrawRectangle(int32(math.Round(float64(p.position.X))), int32(math.Round(float64(p.position.Y))), p.width, p.height, p.color)
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (p *Particle) update(window_width, window_height float32) {
	if p.active {
		p.position.X += p.velocity.X
		p.position.Y += p.velocity.Y

		if rand.Intn(2) == 1 {
			p.color.A -= 1
		}

		if p.color.A == 0 {
			p.active = false
		}
		if p.position.X < 0 || p.position.X+float32(p.width) > window_width {
			p.active = false
		}
		if p.position.Y < 0 || p.position.Y+float32(p.height) > window_height {
			p.active = false
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------
