package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	// Global constants used all over the place.
	screen_width  = int32(1280)
	screen_height = int32(800)
	max_fps       = int32(144)
	title         = "PONNGERS"
)

var (
	// Score tracking
	player_score   = 0
	computer_score = 0
)

// --------------------------------------------------------------------------------------------------------------------

func main() {

	// Window attributes
	rl.InitWindow(screen_width, screen_height, title)
	rl.SetTargetFPS(max_fps)

	// Create objects
	ball := Ball{x_pos: screen_width / 2, y_pos: screen_height / 2, speed_x: 5, speed_y: 5, radius: 20}

	const cpu_paddle_speed = 3 // cpu playes perfectly. I dont think it can miss above 4 and ball speed at 5.
	const player_paddle_speed = 6
	player_paddle := create_paddle(10, true, player_paddle_speed)
	cpu_paddle := create_paddle(rl.GetScreenWidth()-35, false, cpu_paddle_speed)

	// Mainloop.
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// Check for collisions
		if rl.CheckCollisionCircleRec(rl.Vector2{X: float32(ball.x_pos), Y: float32(ball.y_pos)}, ball.radius,
			rl.Rectangle{X: float32(player_paddle.x_pos), Y: float32(player_paddle.y_pos), Width: float32(player_paddle.width), Height: float32(player_paddle.height)}) {
			ball.speed_x *= -1
			ball.x_pos += 5
		}
		if rl.CheckCollisionCircleRec(rl.Vector2{X: float32(ball.x_pos), Y: float32(ball.y_pos)}, ball.radius,
			rl.Rectangle{X: float32(cpu_paddle.x_pos), Y: float32(cpu_paddle.y_pos), Width: float32(cpu_paddle.width), Height: float32(cpu_paddle.height)}) {
			ball.speed_x *= -1
			ball.x_pos -= 5
		}

		// Background
		rl.ClearBackground(rl.NewColor(20, 160, 133, 255))
		rl.DrawRectangle(screen_width/2, 0, screen_width/2, screen_height, rl.NewColor(38, 185, 154, 255))
		rl.DrawCircle(screen_width/2, screen_height/2, 150, rl.NewColor(129, 204, 184, 255))
		rl.DrawCircle(screen_width/2, screen_height/2, 10, rl.White)
		rl.DrawLine(screen_width/2, 0, screen_width/2, screen_height, rl.White)

		// Draw and update the ball.
		ball.Update()
		ball.Draw()

		// Draw and update the paddles.
		player_paddle.Draw()
		player_paddle.Update(int(ball.y_pos))
		cpu_paddle.Draw()
		cpu_paddle.Update(int(ball.y_pos))

		// Scoreboard.
		rl.DrawText(fmt.Sprintf("%v", computer_score), screen_width/4, 20, 100, rl.White)
		rl.DrawText(fmt.Sprintf("%v", player_score), 3*(screen_width/4), 20, 100, rl.White)

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

// --------------------------------------------------------------------------------------------------------------------

type Ball struct {
	x_pos   int32
	y_pos   int32
	speed_x int32
	speed_y int32
	radius  float32
}

func (ball *Ball) Draw() {
	// Draw the ball on the screeen.

	rl.DrawCircle(ball.x_pos, ball.y_pos, ball.radius, rl.Yellow)
}

func (ball *Ball) Update() {
	// Update the position of the ball.

	ball.x_pos += ball.speed_x
	ball.y_pos += ball.speed_y

	// Keep the ball in bounds on the Y axis.
	if int(ball.y_pos+int32(ball.radius)) >= rl.GetScreenHeight() || ball.y_pos-int32(ball.radius) <= 0 {
		ball.speed_y *= -1
	}

	// If the ball goes out of bounds on the X axis, add a point to the respective player.
	if ball.x_pos-int32(ball.radius) >= int32(rl.GetScreenWidth()) {
		computer_score++
		ball.Reset()
	}
	if ball.x_pos-int32(ball.radius) <= 0 {
		player_score++
		ball.Reset()
	}
}

func (ball *Ball) Reset() {
	// Resets ball position and changes direction when out of bounds (score).

	ball.x_pos = int32(rl.GetScreenWidth()) / 2
	ball.y_pos = int32(rl.GetScreenHeight()) / 2

	// Set ball in a random direction.
	var speed_choices = [2]int{-1, 1}
	ball.speed_x *= int32(speed_choices[rl.GetRandomValue(0, 1)])
	ball.speed_y *= int32(speed_choices[rl.GetRandomValue(0, 1)])
}

// --------------------------------------------------------------------------------------------------------------------

type Paddle struct {
	x_pos         int32
	y_pos         int32
	width         int32
	height        int32
	speed         int32
	player_paddle bool
}

func (paddle *Paddle) Draw() {
	// Draw a paddle on the screen.

	rl.DrawRectangleRounded(rl.Rectangle{X: float32(paddle.x_pos), Y: float32(paddle.y_pos), Width: float32(paddle.width), Height: float32(paddle.height)}, 0.8, 0, rl.White)
}

func (paddle *Paddle) Update(ball_y int) {
	// Moves the paddle.

	if paddle.player_paddle {
		// Move the paddle up and down.
		if rl.IsKeyDown(rl.KeyUp) {
			paddle.y_pos -= paddle.speed
		}
		if rl.IsKeyDown(rl.KeyDown) {
			paddle.y_pos += paddle.speed
		}
	}

	// Computer movement algorithm.
	if !paddle.player_paddle {
		if paddle.y_pos+paddle.height/2 > int32(ball_y) {
			paddle.y_pos -= paddle.speed
		}
		if paddle.y_pos+paddle.height/2 <= int32(ball_y) {
			paddle.y_pos += paddle.speed
		}
	}

	// Keep the paddle in bounds.
	if paddle.y_pos <= 0 {
		paddle.y_pos = 0
	}
	if paddle.y_pos+paddle.height >= int32(rl.GetScreenHeight()) {
		paddle.y_pos = int32(rl.GetScreenHeight()) - paddle.height
	}

}

// --------------------------------------------------------------------------------------------------------------------

func create_paddle(x_position int, player bool, paddle_speed int) Paddle {
	// Creates a paddle.

	const paddle_width = 25
	const paddle_height = 120

	return Paddle{
		x_pos: int32(x_position), y_pos: screen_height/2 - paddle_height/2, width: paddle_width, height: paddle_height,
		speed: int32(paddle_speed), player_paddle: player,
	}
}

// --------------------------------------------------------------------------------------------------------------------
