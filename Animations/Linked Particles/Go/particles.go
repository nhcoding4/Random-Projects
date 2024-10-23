package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Particle struct {
	width               int32
	height              int32
	radius              int32
	x                   int32
	y                   int32
	movementX           int32
	movementY           int32
	powerPushMultiplier float64
	pushX               float64
	pushY               float64
	friction            float64
	mouse               *Mouse
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) init() {
	p.setMovementSpeed()
	p.radius = int32(rand.Intn(10)) + 5
	p.pushX = 0
	p.pushY = 0
	p.setFriction()
	p.setPosition()
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) collision() {
	if p.radius > p.x {
		p.x = 1 + p.radius
		p.movementX *= -1
		p.pushX *= -1
	}
	if p.x+p.radius > p.width {
		p.x = p.width - p.radius - 1
		p.movementX *= -1
		p.pushX *= -1
	}

	if p.radius > p.y {
		p.y = 1 + p.radius
		p.movementY *= -1
		p.pushY *= -1
	}
	if p.y+p.radius > p.height {
		p.y = p.height - p.radius - 1
		p.movementY *= -1
		p.pushY *= -1
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) draw() {
	color := rl.ColorFromHSV(float32(p.x), 1.0, 1.0)
	rl.DrawCircle(p.x, p.y, float32(p.radius), color)
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) movement() {
	p.x += p.movementX + int32(math.Round(p.pushX))
	p.y += p.movementY + int32(math.Round(p.pushY))
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) pushParticle() {
	mx, my := p.mouse.getLocation()
	dx := float64(p.x - *mx)
	dy := float64(p.y - *my)

	distance := math.Hypot(dx, dy)
	if distance < float64(*p.mouse.getRadius()) {
		power := (float64(*p.mouse.getRadius()) / distance) * p.powerPushMultiplier
		angle := math.Atan2(dy, dx)
		p.pushX = math.Cos(angle) * power
		p.pushY = math.Sin(angle) * power
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) setFriction() {
	if p.radius <= 7 {
		p.friction = 0.99
	} else if p.radius > 7 && p.radius <= 10 {
		p.friction = 0.97
	} else {
		p.friction = 0.94
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) setMovementSpeed() {
	p.movementX = int32(rand.Intn(4)) - 2
	p.movementY = int32(rand.Intn(4)) - 2
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) setPosition() {
	p.x = int32(rand.Intn(int(p.width)-int(p.radius))) + p.radius
	p.y = int32(rand.Intn(int(p.height)-int(p.radius))) + p.radius
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) update() {
	if *p.mouse.activationStatus() {
		p.pushParticle()
	}
	p.movement()
	p.collision()
	p.updatePush()
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) updatePush() {
	p.pushX *= p.friction
	p.pushY *= p.friction
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) updateWindowSize() {
	p.width = int32(rl.GetScreenWidth())
	p.height = int32(rl.GetScreenHeight())
}

// ---------------------------------------------------------------------------------------------------------------------
