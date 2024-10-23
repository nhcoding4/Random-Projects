package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// ---------------------------------------------------------------------------------------------------------------------

type Simulation struct {
	rows      int32
	columns   int32
	cellSize  int32
	targetFPS int32
	title     string
	fontSize  int32
	grid      *Grid
	gridData  *[][]int32
	ch        chan *[][]int32
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *Simulation) Init() {
	s.ch = make(chan *[][]int32, s.targetFPS)
	s.makeGrid()
	s.windowSetup()
	go s.grid.calculateAndSend()
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *Simulation) drawGrid() {
	s.gridData = <-s.ch

	for i, row := range *s.gridData {
		for j, cell := range row {
			var color rl.Color
			if cell == 0 {
				color = rl.Black
			} else {
				color = rl.White
			}
			rl.DrawRectangle(
				int32(j)*s.cellSize+1,
				int32(i)*s.cellSize+1,
				s.cellSize-1,
				s.cellSize-1,
				color,
			)
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *Simulation) fps() {
	rl.DrawText(fmt.Sprintf("%v", rl.GetFPS()), 0, 0, s.fontSize, rl.Green)
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *Simulation) makeGrid() {
	s.grid = &Grid{
		rows:        s.rows,
		columns:     s.columns,
		grid:        nil,
		ch:          s.ch,
		channelSize: s.targetFPS,
	}
	s.grid.Init()
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *Simulation) run() {
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		s.drawGrid()
		s.fps()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

// ---------------------------------------------------------------------------------------------------------------------

// Use before any drawing function.

func (s *Simulation) windowSetup() {
	width := s.columns * s.cellSize
	height := s.rows * s.cellSize

	rl.InitWindow(width, height, s.title)
	rl.SetTargetFPS(s.targetFPS)
}

// ---------------------------------------------------------------------------------------------------------------------
