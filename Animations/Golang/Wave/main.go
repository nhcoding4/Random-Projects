package main

import (
	"fmt"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Config struct {
	width, height, target_fps int32
	title, current_fps        string
	particle_color            rl.Color
}

type Vector struct {
	X, Y float64
}

type Particle struct {
	radius, deflation_rate float64
	position               Vector
	color                  rl.Color
	active                 bool
}

type Mouse struct {
	position                                 Vector
	radius, angle, angle_speed, x_div, y_div float64
}

func main() {
	// ------------------------------------------------------------------------------------------------------------

	// Init

	// ------------------------------------------------------------------------------------------------------------

	// Config
	var config = Config{
		width:          1000,
		height:         1000,
		target_fps:     144,
		title:          "Bubble Wave",
		particle_color: rl.DarkPurple,
	}

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.SetConfigFlags(rl.FlagWindowHighdpi)
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(config.width, config.height, config.title)
	rl.SetTargetFPS(config.target_fps)

	// Particles
	var particles = []Particle{}

	// Mouse
	var mouse = Mouse{
		position: Vector{
			X: 0.0,
			Y: 0.0,
		},
		radius:      150.0,
		angle:       0.0,
		angle_speed: 2 + rand.Float64(),
		x_div:       540.0,
		y_div:       180.0,
	}

	for !rl.WindowShouldClose() {
		// ------------------------------------------------------------------------------------------------------------

		// Updates

		// ------------------------------------------------------------------------------------------------------------

		// Resize window
		if rl.IsWindowResized() {
			config.width = int32(rl.GetScreenWidth())
			config.height = int32(rl.GetScreenHeight())
		}

		// Mouse
		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			mouse.position.X = float64(rl.GetMouseX())
			mouse.position.Y = float64(rl.GetMouseY())

		} else {
			mouse.angle += mouse.angle_speed
			mouse.position.X = (float64(config.width)-mouse.radius)/2*math.Sin(mouse.angle*math.Pi/mouse.x_div) + float64(config.width)/2
			mouse.position.Y = (float64(config.height)-mouse.radius)/2*math.Cos(mouse.angle*math.Pi/mouse.y_div) + float64(config.height)/2

			if mouse.position.X-mouse.radius < 0 {
				mouse.position.X = mouse.radius
			}
			if mouse.position.X+mouse.radius > float64(config.width) {
				mouse.position.X = float64(config.width) - mouse.radius
			}
			if mouse.position.Y-mouse.radius < 0 {
				mouse.position.Y = mouse.radius
			}
			if mouse.position.Y+mouse.radius > float64(config.height) {
				mouse.position.Y = float64(config.height) - mouse.radius
			}
		}

		// Particles
		var reused = false
		for i := range particles {
			if !particles[i].active {
				particles[i].active = true
				particles[i].position.X = mouse.position.X + float64(float64(rand.Intn(int(mouse.radius))-int(mouse.radius)/2))
				particles[i].position.Y = mouse.position.Y + float64(float64(rand.Intn(int(mouse.radius))-int(mouse.radius)/2))
				particles[i].radius = float64(rand.Intn(40) + 10)
				reused = true
				break
			}
		}

		if !reused {
			var particle = Particle{
				active: true,
				radius: float64(rand.Intn(40) + 10),
				position: Vector{
					X: mouse.position.X + float64(float64(rand.Intn(int(mouse.radius))-int(mouse.radius)/2)),
					Y: mouse.position.Y + float64(float64(rand.Intn(int(mouse.radius))-int(mouse.radius)/2)),
				},
				color:          config.particle_color,
				deflation_rate: 0.1,
			}
			particles = append(particles, particle)
		}

		for i := range particles {
			if particles[i].active {
				particles[i].radius -= particles[i].deflation_rate
				particles[i].color.A = uint8(math.Round(255 / (float64(config.height) / particles[i].position.Y)))

				if particles[i].radius <= 0 {
					particles[i].active = false
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

		// Particles
		for _, particle := range particles {
			if particle.active {
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
		}

		// Fps
		rl.DrawText(config.current_fps, 0, 0, 40, rl.Green)

		rl.EndDrawing()

		// ------------------------------------------------------------------------------------------------------------

	}
	rl.CloseWindow()
}
