package main

import (
	"ago-launcher/config"
	"ago-launcher/gui"
	"ago-launcher/updater"
)

func main() {
	updater := &updater.Updater{}
	updater.GetCurrentModVersion()

	configurator := &config.Configurator{}
	configFile := configurator.LoadConfigFile()
	configurator.ParseConfig(configFile)

	gui.InitGUI(updater, configurator)
}
