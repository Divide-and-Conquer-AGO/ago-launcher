package main

import (
	"runtime/debug"

	"ago-launcher/config"
	"ago-launcher/gui"
	"ago-launcher/news"
	"ago-launcher/quotes"
	"ago-launcher/updater"
	"ago-launcher/utils"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	// Crash Handler
	defer func() {
		if r := recover(); r != nil {
			stack := debug.Stack()
			utils.Logger().Fatalf("error: %v\n%s\n", r, stack)
		}
	}()

	app := app.NewWithID("divide.and.conquer.ago")
	// var splashWindow fyne.Window

	logo := canvas.NewImageFromResource(gui.ResourceFaviconIco)
	logo.FillMode = canvas.ImageFillOriginal
	// logoContainer := container.NewCenter(logo)

	// if drv, ok := app.Driver().(desktop.Driver); ok {
	// 	// splashWindow = drv.CreateSplashWindow()
	// 	splashWindow.SetContent(logoContainer)
	// 	splashWindow.SetTitle("AGO Launcher")
	// 	splashWindow.Show()
	// }

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

	gui.InitGUI(app, nil, updater, configurator, quoter, newsReader)
}
