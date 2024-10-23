package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Projectile struct {
	position, velocity, end rl.Vector2
	width, height           int32
	color                   rl.Color
	active                  bool
}

// --------------------------------------------------------------------------------------------------------------------

func (p *Projectile) init(config *Config, position *rl.Vector2) {
	var calculate_movement = func() rl.Vector2 {
		var dx = p.end.X - p.position.X
		var dy = p.end.Y - p.position.Y

		var movement = rl.Vector2{
			X: dx * 1 / 144.0,
			Y: dy * 1 / 144.0,
		}
		return movement
	}

	p.position = rl.Vector2{
		X: float32(config.width) / 2,
		Y: float32(config.height),
	}
	p.end = *position
	p.velocity = calculate_movement()
	p.width = 2
	p.height = 2
	p.color = rl.White
	p.active = true
}

// --------------------------------------------------------------------------------------------------------------------

func create_projectile(config *Config, position *rl.Vector2) Projectile {
	var new_projectile = Projectile{}
	new_projectile.init(config, position)

	return new_projectile
}

// --------------------------------------------------------------------------------------------------------------------

func (p *Projectile) draw() {
	if p.active {
		rl.DrawRectangle(int32(p.position.X), int32(p.position.Y), p.width, p.height, p.color)
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (p *Projectile) update() bool {
	if p.active {
		p.position.X += p.velocity.X
		p.position.Y += p.velocity.Y

		var dx = p.end.X - p.position.X
		var dy = p.end.Y - p.position.Y
		var distance = math.Hypot(float64(dx), float64(dy))

		if distance < 5 {
			p.active = false
			return true
		}
		return false
	}
	return false
}

// --------------------------------------------------------------------------------------------------------------------
