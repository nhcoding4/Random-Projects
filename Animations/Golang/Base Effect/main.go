package main

import (
	"fmt"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Config struct {
	width, height, target_fps, total_particles int32
	title, current_fps                         string
	particle_color, connector_color            rl.Color
	connection_distance, connection_thickness  float64
}

type Vector struct {
	X, Y float64
}

type Particle struct {
	friction, radius         float64
	position, movement, push Vector
	color                    rl.Color
}

type Mouse struct {
	position Vector
	radius   float64
}

type Connector struct {
	start, end rl.Vector2
	color      rl.Color
}

func main() {
	// ------------------------------------------------------------------------------------------------------------

	// Init

	// ------------------------------------------------------------------------------------------------------------

	// Config
	var config = Config{
		width:                700,
		height:               700,
		target_fps:           144,
		total_particles:      1000,
		connection_distance:  100.0,
		connection_thickness: 3.0,
		title:                "Go Base",
		particle_color:       rl.DarkPurple,
		connector_color:      rl.Magenta,
	}

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.SetConfigFlags(rl.FlagWindowHighdpi)
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(config.width, config.height, config.title)
	rl.SetTargetFPS(config.target_fps)

	// Particles
	var particles = []Particle{}
	for i := 0; i < int(config.total_particles); i++ {

		var radius = float64(rand.Intn(15) + 5)

		var particle = Particle{
			radius: radius,
			position: Vector{
				X: float64(rl.GetRandomValue(int32(radius), config.width-int32(radius))),
				Y: float64(rl.GetRandomValue(int32(radius), config.height-int32(radius))),
			},
			movement: Vector{
				X: float64(rl.GetRandomValue(-1, 0)) + rand.Float64(),
				Y: float64(rl.GetRandomValue(-1, 0)) + rand.Float64(),
			},
			push: Vector{
				X: 0.0,
				Y: 0.0,
			},
			color: config.particle_color,
		}

		if radius < 10 {
			particle.friction = 0.98
		} else if radius > 10 && radius < 15 {
			particle.friction = 0.96
		} else {
			particle.friction = 0.94
		}

		particles = append(particles, particle)
	}

	// Mouse
	var mouse = Mouse{
		position: Vector{
			X: 0.0,
			Y: 0.0,
		},
		radius: 200.0,
	}

	for !rl.WindowShouldClose() {
		// ------------------------------------------------------------------------------------------------------------

		// Updates

		// ------------------------------------------------------------------------------------------------------------

		// Resize window
		if rl.IsWindowResized() {
			config.width = int32(rl.GetScreenWidth())
			config.height = int32(rl.GetScreenHeight())

			for i := range particles {
				particles[i].position = Vector{
					X: float64(rl.GetRandomValue(int32(particles[i].radius), config.width-int32(particles[i].radius))),
					Y: float64(rl.GetRandomValue(int32(particles[i].radius), config.height-int32(particles[i].radius))),
				}
			}
		}

		// Mouse
		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			mouse.position.X = float64(rl.GetMouseX())
			mouse.position.Y = float64(rl.GetMouseY())

			for i := range particles {
				var dx = particles[i].position.X - mouse.position.X
				var dy = particles[i].position.Y - mouse.position.Y
				var distance = math.Hypot(dx, dy)

				if distance <= mouse.radius {
					var power = mouse.radius / distance
					var angle = math.Atan2(dy, dx)
					particles[i].push.X = math.Cos(angle) * power
					particles[i].push.Y = math.Sin(angle) * power
				}
			}
		}

		// Particles
		for i := range particles {
			particles[i].position.X += particles[i].movement.X + particles[i].push.X
			particles[i].position.Y += particles[i].movement.Y + particles[i].push.Y

			particles[i].push.X *= particles[i].friction
			particles[i].push.Y *= particles[i].friction

			if particles[i].position.X-particles[i].radius < 0 {
				particles[i].position.X = particles[i].radius
				particles[i].movement.X *= -1
			}
			if particles[i].position.X+particles[i].radius > float64(config.width) {
				particles[i].position.X = float64(config.width) - particles[i].radius
				particles[i].movement.X *= -1
			}
			if particles[i].position.Y-particles[i].radius < 0 {
				particles[i].position.Y = particles[i].radius
				particles[i].movement.Y *= -1
			}
			if particles[i].position.Y+particles[i].radius > float64(config.height) {
				particles[i].position.Y = float64(config.height) - particles[i].radius
				particles[i].movement.Y *= -1
			}

			particles[i].color.A = uint8(math.Round(255 / (float64(config.height) / particles[i].position.Y)))
		}

		// Connectors
		var connectors = []Connector{}

		for i := range particles {
			for j := i; j < int(config.total_particles); j++ {
				if j == i {
					continue
				}

				var dx = particles[i].position.X - particles[j].position.X
				var dy = particles[i].position.Y - particles[j].position.Y
				var distance = math.Hypot(dx, dy)

				if distance <= config.connection_distance {
					var connector = Connector{
						start: rl.Vector2{
							X: float32(particles[i].position.X),
							Y: float32(particles[i].position.Y),
						},
						end: rl.Vector2{
							X: float32(particles[j].position.X),
							Y: float32(particles[j].position.Y),
						},
						color: config.connector_color,
					}
					connector.color.A = uint8(math.Round(255) * (1 - (distance / config.connection_distance)))

					connectors = append(connectors, connector)
				}
			}
		}

		// Fps
		config.current_fps = fmt.Sprintf("%v", rl.GetFPS())

		// ------------------------------------------------------------------------------------------------------------

		// Drawing

		// ------------------------------------------------------------------------------------------------------------

		rl.BeginDrawing()

		// Background
		rl.ClearBackground(rl.Black)
		rl.DrawRectangleGradientV(0, 0, config.width, config.height, rl.DarkGray, rl.Black)

		// Connectors
		for _, connector := range connectors {
			rl.DrawLineEx(connector.start, connector.end, float32(config.connection_thickness), connector.color)
		}

		// Particles
		for _, particle := range particles {
			rl.DrawCircle(int32(math.Round(particle.position.X)), int32(math.Round(particle.position.Y)), float32(particle.radius+2), rl.Black)
			rl.DrawCircle(int32(math.Round(particle.position.X)), int32(math.Round(particle.position.Y)), float32(particle.radius), rl.White)
			rl.DrawCircle(int32(math.Round(particle.position.X)), int32(math.Round(particle.position.Y)), float32(particle.radius), particle.color)
			rl.DrawCircle(
				int32(math.Round(particle.position.X-particle.radius*0.2)),
				int32(math.Round(particle.position.Y-particle.radius*0.3)),
				float32(particle.radius*0.6),
				rl.White,
			)
		}

		// Fps
		rl.DrawText(config.current_fps, 0, 0, 40, rl.Green)

		rl.EndDrawing()

		// ------------------------------------------------------------------------------------------------------------

	}
	rl.CloseWindow()
}
