package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("AGO Launcher")
	w := a.NewWindow("AGO Launcher")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
    		hello,
    		widget.NewButton("Hi!", func() {
    			hello.SetText("Welcome :)")
    		}),
    	))
	w.ShowAndRun()
}