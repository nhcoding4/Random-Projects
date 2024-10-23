package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

// --------------------------------------------------------------------------------------------------------------------

type app_layout struct{}

// --------------------------------------------------------------------------------------------------------------------

// Order the containers.

func (a *app_layout) Layout() *fyne.Container {
	
	gui := &ui{}
	entry := gui.entry_boxes()
	score := gui.display_score()
	title := gui.title()

	ordered_widgets := container.New(layout.NewVBoxLayout(), title, entry, score)
	padding := container.New(layout.NewPaddedLayout(), ordered_widgets)

	background_rect := gui.background()

	return container.New(layout.NewStackLayout(), background_rect, padding)
}

// --------------------------------------------------------------------------------------------------------------------
