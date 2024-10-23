package main

import (
	"fmt"

	"calculate"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	display_value string = "0"
)

// --------------------------------------------------------------------------------------------------------------------

func main() {

	// Create the Assets.

	// Window
	const width = 600
	const height = 400
	const title = "Love Calculator"
	const fps = 60
	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(fps)

	// Textbox
	textbox_left := textbox([]int{10, 100}, []int{20, 10}, []int{200, 35})
	textbox_right := textbox([]int{rl.GetScreenWidth() - 210, 100}, []int{20, 10}, []int{200, 35})

	// ----------------------------------------------------

	// Main loop: Drawing and updating program state.
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		
		rl.ClearBackground(rl.White)

		
		rl.DrawText("Love Calculator", int32(rl.GetScreenWidth())/3-23, 20, 30, rl.Blue)		
		rl.DrawText(fmt.Sprintf("%v%%", display_value), int32(rl.GetScreenWidth())/2-65, int32(rl.GetScreenHeight())/2, 100, rl.Red)
		
		textbox_left.Draw()
		textbox_right.Draw()

		rl.EndDrawing()
		
		if textbox_left.changed || textbox_right.changed {
			display_value = calculate.Calculate(textbox_left.user_input, textbox_right.user_input)
		}
		textbox_left.changed = false
		textbox_right.changed = false
		if len(textbox_left.user_input) <= 1 || len(textbox_right.user_input) <= 1 {
			display_value = "0"
		}
	}

	// ----------------------------------------------------

	rl.CloseWindow()
}

// --------------------------------------------------------------------------------------------------------------------

func textbox(position, inner_margin, size []int) Textbox {
	new_box := Textbox{
		position:                position,
		inner_margin:            inner_margin,
		size:                    size,
		placeholder:             "Type name here..",
		font_size:               20,
		string_stripper_counter: 0,
		active:                  false,
		normal_colour:           rl.LightGray,
		active_color:            rl.Blue,
		text_colour:             rl.Black,
		placeholder_color:       rl.Gray,
	}
	new_box.Init()
	return new_box
}

// --------------------------------------------------------------------------------------------------------------------
