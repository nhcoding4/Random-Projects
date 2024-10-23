package main

import (
	Deque "dequeDS"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
)

// ---------------------------------------------------------------------------------------------------------------------

// Snake struct and methods which deals with actions related to the snake object.

// ---------------------------------------------------------------------------------------------------------------------

type snake struct {
	addSegment                  bool
	body                        Deque.Deque
	preHeadX, prevHeadY, offset float32
	cellSize                    int32
	direction                   rl.Vector2
	snakeHead                   *rl.Vector2
}

// ---------------------------------------------------------------------------------------------------------------------

// Checks the X direction of the snake for valid modes.

func (s *snake) checkX(direction float32) bool {
	if s.snakeHead.Y == s.prevHeadY {
		return false
	}
	if s.direction.X == direction {
		return false
	}
	return true
}

// --------------------------------------------------------------------------------------------------------------------

// Checks the Y direction of the snake for valid moves.

func (s *snake) checkY(direction float32) bool {
	if s.snakeHead.X == s.preHeadX {
		return false
	}
	if s.direction.Y == direction {
		return false
	}
	return true
}

// --------------------------------------------------------------------------------------------------------------------

// Draws the segments that make up the snake.

func (s *snake) draw() {
	for i := range s.body.Size() {
		segmentData := &s.body.Elems[i]

		floatCellSize := float32(s.cellSize)
		x := segmentData.X * floatCellSize
		y := segmentData.Y * floatCellSize
		segment := rl.Rectangle{
			X:      x + s.offset,
			Y:      y + s.offset,
			Width:  floatCellSize,
			Height: floatCellSize,
		}

		rl.DrawRectangleRounded(segment, 0.5, 6, rl.White)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Functions that need to be executed on object creation.

func (s *snake) init() {
	s.body.Append(rl.Vector2{X: 6, Y: 9}, rl.Vector2{X: 5, Y: 9}, rl.Vector2{X: 4, Y: 9})
	s.direction = rl.Vector2{X: 1, Y: 0}
	s.snakeHead = &s.body.Elems[0]
}

// ---------------------------------------------------------------------------------------------------------------------

// Controls the direction the snake moves.

func (s *snake) moveSnake(xOffset, yOffset float32) {
	s.direction.X = 0 + xOffset
	s.direction.Y = 0 + yOffset
}

// ---------------------------------------------------------------------------------------------------------------------

// Resets the position of the snake back to its starting position.

func (s *snake) resetSnake() {
	s.body.Clear()
	s.init()
}

// --------------------------------------------------------------------------------------------------------------------

// Updates the position of the segments.

func (s *snake) update() {
	newSegment := rl.Vector2Add(*s.snakeHead, s.direction)
	s.body.Prepend(newSegment)
	s.snakeHead = &newSegment

	if s.addSegment {
		s.addSegment = false
	} else {
		_, err := s.body.RemoveBack(1)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Updates the previous head position held in the struct.

func (s *snake) updatePrevHeadPosition() {
	s.preHeadX = s.snakeHead.X
	s.prevHeadY = s.snakeHead.Y
}

// --------------------------------------------------------------------------------------------------------------------
