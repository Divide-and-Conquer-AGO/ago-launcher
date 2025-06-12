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
	configurator.ConfigLocations = []string{".", "config", "eopData/config"}
	configurator.LoadAllConfigFiles()

	// NEWSREADER
	newsReader := &news.NewsReader{}
	newsReader.GetNewsItems()

	gui.InitGUI(updater, configurator, quoter, newsReader)
}
