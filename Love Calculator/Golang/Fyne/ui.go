package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"calculate"
)

// --------------------------------------------------------------------------------------------------------------------

type ui struct {
	score         *canvas.Text
	entry_widgets []*widget.Entry
}

// --------------------------------------------------------------------------------------------------------------------

// Entry boxes and button. But can be seperated by deleting the button call and placed seperately in the Ui container.

func (u *ui) entry_boxes() *fyne.Container {

	for i := 0; i < 2; i++ {
		new_entry := widget.NewEntry()
		new_entry.SetPlaceHolder("Enter a name...")
		u.entry_widgets = append(u.entry_widgets, new_entry)
	}
	button := u.button()

	sub_container := container.New(layout.NewGridLayoutWithColumns(3), u.entry_widgets[0], button, u.entry_widgets[1])
	return container.New(layout.NewPaddedLayout(), sub_container)
}

// --------------------------------------------------------------------------------------------------------------------

// Score label.

func (u *ui) display_score() *fyne.Container {

	u.score = canvas.NewText("0%", color.White)
	u.score.TextSize = 50
	u.score.Alignment = fyne.TextAlignCenter

	sub_container := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), u.score, layout.NewSpacer())
	return container.New(layout.NewPaddedLayout(), sub_container)
}

// --------------------------------------------------------------------------------------------------------------------

// Button, proccesses user input.

func (u *ui) button() *fyne.Container {

	button := widget.NewButton("Calculate", func() {
		u.score.Text = calculate.Calculate(u.entry_widgets[0].Text, u.entry_widgets[1].Text) + "%"
		u.score.Refresh()
	})
	sub_container := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	return container.New(layout.NewPaddedLayout(), sub_container)
}

// --------------------------------------------------------------------------------------------------------------------

// Background rectangle.

func (u *ui) background() *fyne.Container {

	rect := canvas.NewRectangle(color.RGBA{R: 25, G: 25, B: 25, A: 240})
	return container.New(layout.NewPaddedLayout(), rect)
}

// --------------------------------------------------------------------------------------------------------------------

// Title.

func (u *ui) title() *fyne.Container {

	title := canvas.NewText("Love Calculator", color.White)
	title.TextSize = 30
	title.Alignment = fyne.TextAlignCenter

	return container.New(layout.NewPaddedLayout(), title)
}
