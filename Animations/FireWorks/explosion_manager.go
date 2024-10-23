package main

import (
	"math/rand"
	"time"
)

type Explosion_Manager struct {
	mouse       Mouse
	particles   []Particle
	projectiles []Projectile
	last_fired  time.Time
	time_delta  int32
}

// --------------------------------------------------------------------------------------------------------------------

func (e *Explosion_Manager) init(config *Config) {
	e.mouse = create_mouse(config)
	e.time_delta = 3
	e.fire_projectile(config)
	e.last_fired = time.Now()
}

// --------------------------------------------------------------------------------------------------------------------

func (e *Explosion_Manager) create_explosion(config *Config, index int) {
	e.mouse.update_position(config)

	var particles_in_effect = rand.Intn(1000) + 50

	for i := range e.particles {
		if !e.particles[i].active {
			e.particles[i].init(&e.projectiles[index].end)
			particles_in_effect--
		}

		if particles_in_effect == 0 {
			break
		}
	}

	for particles_in_effect > 0 {
		e.particles = append(e.particles, create_particle(&e.projectiles[index].end))
		particles_in_effect--
	}

}

// --------------------------------------------------------------------------------------------------------------------

func (e *Explosion_Manager) draw() {
	for i := range e.particles {
		e.particles[i].draw()
	}

	for i := range e.projectiles {
		e.projectiles[i].draw()
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (e *Explosion_Manager) fire_projectile(config *Config) {
	e.mouse.update_position(config)

	var fire = func() {
		for i := range e.projectiles {
			if !e.projectiles[i].active {
				e.projectiles[i].init(config, &e.mouse.position)
				e.last_fired = time.Now()
				return
			}
		}

		e.projectiles = append(e.projectiles, create_projectile(config, &e.mouse.position))

		e.last_fired = time.Now()
	}

	if e.mouse.clicked {
		fire()
		e.mouse.clicked = false
		return
	}

	if time.Since(e.last_fired) >= time.Duration(e.time_delta)*time.Second {
		fire()
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (e *Explosion_Manager) update(config *Config) {
	for i := range e.particles {
		e.particles[i].update(float32(config.width), float32(config.height))
	}

	for i := range e.projectiles {
		if e.projectiles[i].update() {
			e.create_explosion(config, i)
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------
