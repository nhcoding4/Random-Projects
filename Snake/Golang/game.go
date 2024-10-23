package main

import (
	FileActions "fileActions"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"os"
)

// ---------------------------------------------------------------------------------------------------------------------

// Game struct and methods which controls all other objects and their relationship with one another.

// ---------------------------------------------------------------------------------------------------------------------

type game struct {
	running                                          bool
	backgroundColour                                 color.RGBA
	borderOffset                                     float32
	lastUpdate, updateInterval                       float64
	gameFood                                         *food
	targetFPS, cellSize, cellCount, score, highScore int32
	snake                                            *snake
	wallSound, eatSound                              *rl.Sound
	title                                            string
}

// ---------------------------------------------------------------------------------------------------------------------

// Changes the speed of the game.

func (g *game) changeSpeed() {
	if g.score%2 == 0 {
		g.updateInterval *= 0.95
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Draws border around screen.

func (g *game) drawBorder() {
	rl.DrawRectangleLinesEx(
		rl.Rectangle{
			X:      g.borderOffset - 5,
			Y:      g.borderOffset - 5,
			Width:  float32(g.cellSize*g.cellCount) + 10,
			Height: float32(g.cellSize*g.cellCount) + 10,
		},
		5,
		rl.White,
	)
}

// ---------------------------------------------------------------------------------------------------------------------

// Checks for collision with food.

func (g *game) checkFoodCollision() {
	if rl.Vector2Equals(*g.snake.snakeHead, g.gameFood.foodPosition) {
		g.gameFood.generateRandomPosition(&g.snake.body.Elems)
		g.snake.addSegment = true
		g.score++
		g.changeSpeed()
		rl.PlaySound(*g.eatSound)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Checks if the snake head collides with the edges or itself.

func (g *game) checkSnakeCollision() {
	if int32(g.snake.snakeHead.X) == g.cellCount || g.snake.snakeHead.X == -1 {
		g.gameOver()
		rl.PlaySound(*g.wallSound)
	}
	if int32(g.snake.snakeHead.Y) == g.cellCount || g.snake.snakeHead.Y == -1 {
		g.gameOver()
		rl.PlaySound(*g.wallSound)
	}

	for i := 0; i < g.snake.body.Size(); i++ {
		if i == 0 {
			continue
		}
		if g.snake.snakeHead.X == g.snake.body.Elems[i].X && g.snake.snakeHead.Y == g.snake.body.Elems[i].Y {
			g.gameOver()
			rl.PlaySound(*g.wallSound)
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Listens for key presses and changes snake direction.

func (g *game) checkUserInput() {
	const up = -1
	const down = 1
	const left = -1
	const right = 1

	if rl.IsKeyPressed(rl.KeyUp) && g.snake.checkY(down) {
		g.snake.moveSnake(0, -1)
		g.snake.updatePrevHeadPosition()
	}
	if rl.IsKeyPressed(rl.KeyDown) && g.snake.checkY(up) {
		g.snake.moveSnake(0, 1)
		g.snake.updatePrevHeadPosition()
	}
	if rl.IsKeyPressed(rl.KeyLeft) && g.snake.checkX(right) {
		g.snake.moveSnake(-1, 0)
		g.snake.updatePrevHeadPosition()
	}
	if rl.IsKeyPressed(rl.KeyRight) && g.snake.checkX(left) {
		g.snake.moveSnake(1, 0)
		g.snake.updatePrevHeadPosition()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Draws game elements on the screen.

func (g *game) draw() {
	g.drawBorder()
	g.text()
	g.gameFood.draw()
	g.snake.draw()
}

// ---------------------------------------------------------------------------------------------------------------------

// Checks if X amount of time has passed.

func (g *game) eventTriggered() bool {
	currentTime := rl.GetTime()
	if currentTime-g.lastUpdate >= g.updateInterval {
		g.lastUpdate = currentTime
		return true
	}
	return false
}

// ---------------------------------------------------------------------------------------------------------------------

// Quit game actions.

func (g *game) exitActions() {
	const folderName = "temp"
	err := os.RemoveAll(folderName)
	if err != nil {
		fmt.Println(err)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// End of game actions.

func (g *game) gameOver() {
	g.snake.resetSnake()
	g.gameFood.generateRandomPosition(&g.snake.body.Elems)
	g.running = false

	if g.score > g.highScore {
		directoryLocation, highScoreLocation, err := filePaths()
		if err != nil {
			fmt.Println(err)
		}
		g.highScore = g.score
		err = FileActions.WriteScore(highScoreLocation, directoryLocation, int(g.score))
		if err != nil {
			fmt.Println(err)
		}
	}

	g.updateInterval = 0.2
	g.score = 0
}

// ---------------------------------------------------------------------------------------------------------------------

// Deals with initialising all required game objects.

func (g *game) init() {
	rl.InitWindow(
		int32(2*g.borderOffset)+(g.cellSize*g.cellCount),
		int32(2*g.borderOffset)+(g.cellSize*g.cellCount),
		g.title,
	)
	rl.SetTargetFPS(g.targetFPS)

	rl.InitAudioDevice()
	g.loadSounds()

	g.snake = &snake{cellSize: g.cellSize, offset: g.borderOffset}
	g.snake.init()

	g.gameFood = &food{
		cellSize:  g.cellSize,
		cellCount: g.cellCount,
		color:     rl.White,
		offset:    int32(g.borderOffset),
	}
	g.gameFood.init(&g.snake.body.Elems)

	_, highScoreLocation, err := filePaths()
	if err != nil {
		fmt.Println(err)
	}
	highScore, err := FileActions.LoadScore(highScoreLocation)
	if err != nil {
		fmt.Println(err)
	}
	g.highScore = int32(highScore)
}

// ---------------------------------------------------------------------------------------------------------------------

// Loads the game sounds.

func (g *game) loadSounds() {
	const wallSoundLocation = "temp/wall.mp3"
	const eatSoundLocation = "temp/eat.mp3"

	wallSound := rl.LoadSound(wallSoundLocation)
	eatSound := rl.LoadSound(eatSoundLocation)
	g.wallSound = &wallSound
	g.eatSound = &eatSound
}

// ---------------------------------------------------------------------------------------------------------------------

// Starts and stops the game state.

func (g *game) pauseUnpause() {
	if rl.IsKeyPressed(rl.KeySpace) {
		g.running = !g.running
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Draws text onto the screen.

func (g *game) text() {
	highScoreMessage := fmt.Sprintf("High Score: %v", g.highScore)
	scoreMessage := fmt.Sprintf("Score: %v", g.score)

	rl.DrawText(highScoreMessage, int32(g.borderOffset+5), 20, 40, rl.White)
	rl.DrawText(scoreMessage, int32(g.borderOffset+575), 20, 40, rl.White)
}

// ---------------------------------------------------------------------------------------------------------------------

// Updates the state of all game object.

func (g *game) update() {
	g.pauseUnpause()
	if g.running {
		g.checkUserInput()
		g.checkFoodCollision()
		g.checkSnakeCollision()

		if g.eventTriggered() {
			g.snake.update()
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------
