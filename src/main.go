package main

import (
	"fmt"

	"ago-launcher/src/gui"
)

func main() {
	fmt.Println("Launching AGO launcher..")

	body := gui.InitGUI()
	gui.RenderInputs(body)
	body.RunMainWindow()
}
