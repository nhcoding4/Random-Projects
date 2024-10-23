package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// --------------------------------------------------------------------------------------------------------------------

type Textbox struct {
	active                  bool
	changed                 bool
	position                []int
	inner_margin            []int
	font_size               int
	size                    []int
	string_stripper_counter int
	text_colour             rl.Color
	normal_colour           rl.Color
	active_color            rl.Color
	placeholder_color       rl.Color
	background_rectangle    rl.Rectangle
	user_input              string
	placeholder             string
}

// --------------------------------------------------------------------------------------------------------------------

// Sets up the textbox. Needs to be called on struct creation to allow for other functions to work.
func (t *Textbox) Init() {
	t.background_rectangle = rl.NewRectangle(float32(t.position[0]), float32(t.position[1]), float32(t.size[0]), float32(t.size[1]))
}

// --------------------------------------------------------------------------------------------------------------------

// Draw the textbox on the screen.
func (t *Textbox) Draw() {

	width := rl.MeasureText(t.user_input[t.string_stripper_counter:], int32(t.font_size))

	// Set the colour depending on if the textbox is active or not.
	var colour rl.Color
	if t.active {
		colour = t.active_color
	} else {
		colour = t.normal_colour
	}
	rl.DrawRectangleRec(t.background_rectangle, colour)

	// ----------------------------------------------------

	// Prevent text overflow.
	if width+int32(t.inner_margin[0]) >= int32(t.size[0]) {
		t.string_stripper_counter++
	}
	if t.active {
		rl.DrawText(
			t.user_input[t.string_stripper_counter:],
			int32(t.position[0]+t.inner_margin[0]),
			int32(t.position[1]+t.inner_margin[1]),
			int32(t.font_size),
			t.text_colour,
		)
	} else {
		if len(t.user_input) == 0 {
			rl.DrawText(
				t.placeholder,
				int32(t.position[0]+t.inner_margin[0]),
				int32(t.position[1]+t.inner_margin[1]),
				int32(t.font_size),
				t.placeholder_color,
			)
		}
		rl.DrawText(
			t.user_input[:len(t.user_input)-t.string_stripper_counter],
			int32(t.position[0]+t.inner_margin[0]),
			int32(t.position[1]+t.inner_margin[1]),
			int32(t.font_size),
			t.text_colour,
		)
	}
	t.Input()
}

// --------------------------------------------------------------------------------------------------------------------

// Take input from the user.
func (t *Textbox) Input() {
	// Check if the user's mouse is over the textbox.
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), t.background_rectangle) {
		t.active = true
	} else {
		t.active = false
	}

	if t.active {
		// Take input. Use a endless loop to deal with multiple inputs per frame.
		key := rl.GetCharPressed()
		for key > 0 {
			if key >= 0 && key <= 255 {
				t.user_input += string(key)
			}
			key = rl.GetCharPressed()
		}

		// Delete input.
		if rl.IsKeyPressed(rl.KeyBackspace) && len(t.user_input) > 0 {
			t.user_input = t.user_input[:len(t.user_input)-1]
			if t.string_stripper_counter > 0 {
				t.string_stripper_counter--
			}
		}
		t.changed = true
	}
}

// --------------------------------------------------------------------------------------------------------------------
