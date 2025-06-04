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
	// AGO.cfg
	configurator.AGOConfigFile = configurator.LoadConfigFile("AGO.cfg")
	configurator.ParseConfig(configurator.AGOConfigFile, &configurator.AGOConfig)
	
	// TATW.cfg
	configurator.ModConfigFile = configurator.LoadConfigFile("TATW.cfg")
	configurator.ParseConfig(configurator.ModConfigFile, &configurator.ModConfig)

	gui.InitGUI(updater, configurator)
}
