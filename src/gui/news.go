package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getNewsContent() fyne.CanvasObject {
	// Check for Updates
	newsWidget := widget.NewLabel("Some news item")
	// Container
	content := container.NewVBox(
		newsWidget,
	)

	return content
}
