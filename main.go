package main

import (
	"ago-launcher/gui"
	"fmt"
)

func main() {
	fmt.Println("Launching AGO launcher..")

	body := gui.InitGUI()
	gui.RenderInputs(body)
	body.RunMainWindow()
}
