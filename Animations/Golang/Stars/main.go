package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// --------------------------------------------------------------------------------------------------------------------

// Types

// --------------------------------------------------------------------------------------------------------------------

type Config struct {
	width, height, target_fps, total_particles int32
	title, current_fps                         string
	particle_color, connector_color            rl.Color
	connection_distance, connection_thickness  float64
}

type Sprite struct {
	sprite_texture     rl.Texture2D
	width, height      float32
	x_images, y_images int32
}

type Vector struct {
	X, Y float64
}

type Particle struct {
	friction, width, height, update_delta float64
	position, movement, push              Vector
	color                                 rl.Color
	current_image_x, current_image_y      int32
	image_rectangle                       rl.Rectangle
	last_update                           time.Time
}

type Mouse struct {
	position Vector
	radius   float64
}

type Connector struct {
	start, end rl.Vector2
	color      rl.Color
}

// --------------------------------------------------------------------------------------------------------------------

func main() {
	// ----------------------------------------------------------------------------------------------------------------

	// Init

	// ----------------------------------------------------------------------------------------------------------------

	// Config

	var config = Config{
		width:                1000,
		height:               1000,
		target_fps:           144,
		total_particles:      500,
		connection_distance:  100.0,
		connection_thickness: 2.0,
		title:                "Stars",
		particle_color:       rl.White,
		connector_color:      rl.White,
	}

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.SetConfigFlags(rl.FlagWindowHighdpi)
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(config.width, config.height, config.title)
	rl.SetTargetFPS(config.target_fps)

	// ----------------------------------------------------------------------------------------------------------------

	// Spirtes

	var sprite_image rl.Image = *rl.LoadImage("./stars_sprite.png")

	var sprite = Sprite{
		sprite_texture: rl.LoadTextureFromImage(&sprite_image),
		width:          50,
		height:         50,
		y_images:       3,
		x_images:       3,
	}
	defer rl.UnloadTexture(sprite.sprite_texture)

	// ----------------------------------------------------------------------------------------------------------------

	// Particles

	var particles = []Particle{}
	for i := 0; i < int(config.total_particles); i++ {

		var particle = Particle{
			width:  float64(sprite.width),
			height: float64(sprite.height),
			position: Vector{
				X: float64(rl.GetRandomValue(0, config.width-int32(sprite.width/2))),
				Y: float64(rl.GetRandomValue(0, config.height-int32(sprite.height/2))),
			},
			movement: Vector{
				X: (float64(rl.GetRandomValue(-1, 0)) + rand.Float64()) / 2.0,
				Y: (float64(rl.GetRandomValue(-1, 0)) + rand.Float64()) / 2.0,
			},
			push: Vector{
				X: 0.0,
				Y: 0.0,
			},
			friction:        0.95,
			last_update:     time.Now(),
			update_delta:    1.0 / float64(config.target_fps),
			color:           config.particle_color,
			current_image_x: 0,
			current_image_y: rl.GetRandomValue(0, sprite.y_images-1),
		}
		particles = append(particles, particle)
	}

	// ----------------------------------------------------------------------------------------------------------------

	// Mouse

	var mouse = Mouse{
		position: Vector{
			X: 0.0,
			Y: 0.0,
		},
		radius: 200.0,
	}

	// ----------------------------------------------------------------------------------------------------------------

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
					X: float64(rl.GetRandomValue(0, config.width-int32(sprite.width))),
					Y: float64(rl.GetRandomValue(0, config.height-int32(sprite.height))),
				}
			}
		}

		// ------------------------------------------------------------------------------------------------------------

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

		// ------------------------------------------------------------------------------------------------------------

		// Particles

		for i := range particles {
			// Movement
			particles[i].position.X += particles[i].movement.X + particles[i].push.X
			particles[i].position.Y += particles[i].movement.Y + particles[i].push.Y

			particles[i].push.X *= particles[i].friction
			particles[i].push.Y *= particles[i].friction

			if particles[i].position.X-particles[i].width/2 < 0 {
				particles[i].position.X = particles[i].width / 2
				particles[i].movement.X *= -1
			}
			if particles[i].position.X+particles[i].width/2 > float64(config.width) {
				particles[i].position.X = float64(config.width) - particles[i].width/2
				particles[i].movement.X *= -1
			}
			if particles[i].position.Y-particles[i].width/2 < 0 {
				particles[i].position.Y = particles[i].width / 2
				particles[i].movement.Y *= -1
			}
			if particles[i].position.Y+particles[i].height/2 > float64(config.height) {
				particles[i].position.Y = float64(config.height) - particles[i].height/2
				particles[i].movement.Y *= -1
			}

			// Image update
			if time.Since(particles[i].last_update) > time.Duration(particles[i].update_delta*float64(time.Microsecond)) {
				particles[i].current_image_x++

				if particles[i].current_image_x > sprite.x_images-1 {
					particles[i].current_image_x = 0
				}

				particles[i].image_rectangle = rl.Rectangle{
					X:      float32(particles[i].current_image_x) * sprite.width,
					Y:      float32(particles[i].current_image_y) * sprite.height,
					Width:  sprite.width,
					Height: sprite.height,
				}

				particles[i].last_update = time.Now()
			}

		}

		// ------------------------------------------------------------------------------------------------------------

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

		// ------------------------------------------------------------------------------------------------------------

		// Fps

		config.current_fps = fmt.Sprintf("%v", rl.GetFPS())

		// ------------------------------------------------------------------------------------------------------------

		// Drawing

		// ------------------------------------------------------------------------------------------------------------

		rl.BeginDrawing()

		// Background

		rl.ClearBackground(rl.Black)

		// ------------------------------------------------------------------------------------------------------------

		// Connectors

		for _, connector := range connectors {
			rl.DrawLineEx(connector.start, connector.end, float32(config.connection_thickness), connector.color)
		}

		// ------------------------------------------------------------------------------------------------------------

		// Particles

		for _, particle := range particles {
			rl.DrawTexturePro(
				sprite.sprite_texture,
				particle.image_rectangle,
				rl.Rectangle{
					X:      float32(particle.position.X),
					Y:      float32(particle.position.Y),
					Width:  sprite.width,
					Height: sprite.height,
				},
				rl.Vector2{
					X: float32(sprite.width) / 2,
					Y: float32(sprite.height) / 2,
				},
				0,
				rl.White,
			)
		}

		// ------------------------------------------------------------------------------------------------------------

		// Fps

		rl.DrawText(config.current_fps, 0, 0, 40, rl.Green)

		rl.EndDrawing()

		// ------------------------------------------------------------------------------------------------------------

	}
	rl.CloseWindow()
}
