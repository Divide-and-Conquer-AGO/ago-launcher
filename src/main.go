package main

import (
	"ago-launcher/config"
	"ago-launcher/gui"
	"ago-launcher/news"
	"ago-launcher/quotes"
	"ago-launcher/updater"
)

func main() {
	// QUOTER
	quoter := &quotes.Qouter{}

	// UPDATER
	updater := &updater.Updater{}
	updater.CheckForUpdate()

	// CONFIG
	configurator := &config.Configurator{}
	// AGO.cfg
	configurator.AGOConfigFile = configurator.LoadConfigFile("AGO.cfg")
	configurator.ParseConfig(configurator.AGOConfigFile, &configurator.AGOConfig)
	
	// TATW.cfg
	configurator.ModConfigFile = configurator.LoadConfigFile("TATW.cfg")
	configurator.ParseConfig(configurator.ModConfigFile, &configurator.ModConfig)

	// NEWSREADER
	newsReader := &news.NewsReader{}
	newsReader.GetNewsItems()

	gui.InitGUI(updater, configurator, quoter, newsReader)
}
