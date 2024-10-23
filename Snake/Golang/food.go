package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

// ---------------------------------------------------------------------------------------------------------------------

// Food struct and methods which deals with actions related to the food object.

// ---------------------------------------------------------------------------------------------------------------------

type food struct {
	color                       color.RGBA
	cellSize, cellCount, offset int32
	foodTexture                 rl.Texture2D
	foodPosition                rl.Vector2
}

// ---------------------------------------------------------------------------------------------------------------------

// Draws a representation of the food.

func (f *food) draw() {
	x := int32(f.foodPosition.X)*f.cellSize + f.offset
	y := int32(f.foodPosition.Y)*f.cellSize + f.offset
	rl.DrawTexture(f.foodTexture, x, y, rl.White)
}

// ---------------------------------------------------------------------------------------------------------------------

// Generates a random position for the food to appear at.

func (f *food) generateRandomPosition(snakePos *[]rl.Vector2) {
	for {
		f.foodPosition.X = float32(rl.GetRandomValue(0, f.cellCount-1))
		f.foodPosition.Y = float32(rl.GetRandomValue(0, f.cellCount-1))

		foodOverlap := func() bool {
			for _, segment := range *snakePos {
				if rl.Vector2Equals(f.foodPosition, segment) {
					return true
				}
			}
			return false
		}()

		if !foodOverlap {
			break
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Functions needed to be executed on object creation.

func (f *food) init(snakePos *[]rl.Vector2) {
	f.loadImage()
	f.generateRandomPosition(snakePos)
}

// ---------------------------------------------------------------------------------------------------------------------

// Loads a texture from file.

func (f *food) loadImage() {
	const imageLocation = "temp/food.png"
	image := rl.LoadImage(imageLocation)
	f.foodTexture = rl.LoadTextureFromImage(image)
	rl.UnloadImage(image)
}

// ---------------------------------------------------------------------------------------------------------------------
