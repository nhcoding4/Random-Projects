package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// --------------------------------------------------------------------------------------------------------------------

// The internet time. GMT + 1 converted into beats. (1000 per day.)

type Current_Time struct {
	current time.Time
}

func (t *Current_Time) Settime() {
	location, _ := time.LoadLocation("Europe/London")
	t.current = time.Now().In(location)
}

func (t *Current_Time) String() string {

	hour_beats := float64(t.current.Hour()) * 41.666
	min_beats := float64(t.current.Minute()) * 0.6944
	second_beats := float64(t.current.Second()) * 0.01157

	return fmt.Sprintf("%.3f", hour_beats+min_beats+second_beats)

}

// --------------------------------------------------------------------------------------------------------------------

func main() {

	// Window setup.

	const width = 400
	const height = 150
	const title = "Internet Time"
	const fps = 60

	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(fps)

	internet_time := Current_Time{}

	const width_start = 15
	const height_start = 2
	const font_size = 100

	const unit_text = "beats"
	const unit_font_size = 50
	unit_width := rl.GetScreenWidth()/2 - 75
	unit_height := 90

	// ----------------------------------------------------

	// Mainloop

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		internet_time.Settime()
		rl.DrawText(internet_time.String(), width_start, height_start, font_size, rl.White)
		rl.DrawText(unit_text, int32(unit_width), int32(unit_height), unit_font_size, rl.White)

		rl.EndDrawing()
	}

	// ----------------------------------------------------

	rl.CloseWindow()
}

// --------------------------------------------------------------------------------------------------------------------
