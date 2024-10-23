package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

type Game struct {
	windowWidth         int32
	windowHeight        int32
	cellSize            int32
	targetFPS           int32
	title               string
	rows                int32
	columns             int32
	forcedMovementDelta time.Duration
	colors              []rl.Color
	grid                Grid
	movementTick        time.Time
	downMovement        time.Time
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) init() {
	g.windowHeight = g.rows * g.cellSize
	g.windowWidth = g.columns * g.cellSize
	rl.InitWindow(g.windowWidth, g.windowHeight, g.title)
	rl.SetTargetFPS(g.targetFPS)

	g.loadColors()
	g.createGrid()
	g.forcedMovementDelta = 1 * time.Second
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) autoMoveDown() {
	delta := time.Now().Sub(g.downMovement)

	if delta > g.forcedMovementDelta {
		g.grid.moveBlock("down")
		g.downMovement = time.Now()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) checkNewBlock() {
	for _, cell := range g.grid.currentBlock[g.grid.rotation] {
		y := cell[1]
		if y+g.grid.yOffset == g.rows-1 {
			g.grid.selectBlock()
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) checkCollisionX(direction int32) bool {
	positions := g.grid.blockCoordinates()

	for _, cell := range g.grid.currentBlock[g.grid.rotation] {
		if g.grid.checkSelfX(cell[0], cell[1], direction, positions) {
			continue
		}
		nextX := cell[0] + g.grid.xOffset + direction
		nextY := cell[1] + g.grid.yOffset

		if nextX < 0 || nextX > g.columns-1 {
			return false
		}

		if g.grid.cells[nextY][nextX] != 0 {
			return false
		}
	}
	return true
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) collisionY() {
	positions := g.grid.blockCoordinates()
	newBlock := false

	for _, cell := range g.grid.currentBlock[g.grid.rotation] {
		if g.grid.checkSelfY(cell[0], cell[1], positions) {
			continue
		}

		nextX := cell[0] + g.grid.xOffset
		nextY := cell[1] + g.grid.yOffset + 1

		if nextY > g.rows-1 {
			break
		}

		if g.grid.cells[nextY][nextX] != 0 {
			newBlock = true
		}
	}
	if newBlock {
		g.grid.selectBlock()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) checkValidMoveLeft() bool {
	for _, cell := range g.grid.currentBlock[g.grid.rotation] {
		x := cell[0]
		if x+g.grid.xOffset-1 < 0 {
			return false
		}
		if !g.checkCollisionX(-1) {
			return false
		}
	}
	return true
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) checkValidMoveRight() bool {
	for _, cell := range g.grid.currentBlock[g.grid.rotation] {
		x := cell[0]
		if x+g.grid.xOffset+1 > g.columns-1 {
			return false
		}
		if !g.checkCollisionX(1) {
			return false
		}
	}
	return true
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) checkValidRotation() bool {
	nextRotation := g.grid.rotation + 1
	maxRotations := int32(3)
	if nextRotation > maxRotations {
		nextRotation = int32(0)
	}

	for _, cell := range g.grid.currentBlock[nextRotation] {
		x := cell[0]
		if x+g.grid.xOffset < 0 {
			return false
		}
		if x+g.grid.xOffset > g.columns-1 {
			return false
		}
	}
	return true
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) createGrid() {
	g.grid = Grid{
		cellSize: g.cellSize,
		rows:     g.rows,
		columns:  g.columns,
		colors:   g.colors,
	}
	g.grid.init()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) increaseForcedMovementSpeed() {
	if g.grid.totalLinesCleared >= 1 {
		g.forcedMovementDelta /= 100
		g.forcedMovementDelta *= 90
		g.grid.totalLinesCleared = 0
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) loadColors() {
	g.colors = []rl.Color{
		rl.DarkGray,
		rl.Green,
		rl.Red,
		rl.Orange,
		rl.Yellow,
		rl.Purple,
		rl.SkyBlue,
		rl.Blue,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) run() {
	for !rl.WindowShouldClose() {
		g.collisionY()
		g.checkNewBlock()
		g.userInput()
		g.autoMoveDown()
		g.grid.updateGrid()
		g.increaseForcedMovementSpeed()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		g.grid.drawGrid()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) userInput() {
	delta := time.Now().Sub(g.movementTick)

	if rl.IsKeyDown(rl.KeyLeft) && delta.Seconds() > 0.1 && g.checkValidMoveLeft() {
		g.grid.moveBlock("left")
		g.movementTick = time.Now()
	}
	if rl.IsKeyDown(rl.KeyRight) && delta.Seconds() > 0.1 && g.checkValidMoveRight() {
		g.grid.moveBlock("right")
		g.movementTick = time.Now()
	}
	if rl.IsKeyDown(rl.KeyDown) && delta.Seconds() > 0.1 {
		g.grid.moveBlock("down")
		g.movementTick = time.Now()
	}

	maxRotation := int32(3)
	if rl.IsKeyPressed(rl.KeyUp) && g.checkValidRotation() {
		g.grid.clearPrevious()
		g.grid.rotation++
		if g.grid.rotation > maxRotation {
			g.grid.rotation = 0
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------
