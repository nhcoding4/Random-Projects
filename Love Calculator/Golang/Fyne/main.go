package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// --------------------------------------------------------------------------------------------------------------------

func main() {

	application := app.New()
	window := application.NewWindow("Love Calculator")
	window.Resize(fyne.NewSize(500, 125))

	new_layout := &app_layout{}
	window.SetContent(new_layout.Layout())

	window.ShowAndRun()
}

// --------------------------------------------------------------------------------------------------------------------
