package main

import (
	"ago-launcher/config"
	"ago-launcher/gui"
	"ago-launcher/news"
	"ago-launcher/quotes"
	"ago-launcher/updater"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
)

func main() {
	app := app.NewWithID("divide.and.conquer.ago")

	logo := canvas.NewImageFromResource(gui.ResourceFaviconIco)
	logo.FillMode = canvas.ImageFillOriginal
	logoContainer := container.NewCenter(logo)

	var splashWindow fyne.Window
	if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
        splashWindow = drv.CreateSplashWindow()
        splashWindow.SetContent(logoContainer)
		splashWindow.SetTitle("AGO Launcher")
        splashWindow.Show()
    }

	// QUOTER
	quoter := &quotes.Qouter{}

	// UPDATER
	updater := &updater.Updater{}
	updater.CheckForUpdate()

	// CONFIG
	configurator := &config.Configurator{}
	configurator.LoadAllConfigFiles()

	// NEWSREADER
	newsReader := &news.NewsReader{}
	newsReader.GetNewsItems()

	gui.InitGUI(app, splashWindow, updater, configurator, quoter, newsReader)
}
