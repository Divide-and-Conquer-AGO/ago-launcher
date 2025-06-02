package main

import (
	"ago-launcher/gui"
	"ago-launcher/updater"
)

func main() {
	updater := &updater.Updater{}
	updater.GetCurrentModVersion()

	gui.InitGUI(updater)
}
