package gui

import (
	"ago-launcher/config"
	"ago-launcher/updater"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func InitGUI(updater *updater.Updater, configurator *config.Configurator) {
	myApp := app.NewWithID("divide.and.conquer.ago")

	myWindow := myApp.NewWindow("AGO Launcher")
	// myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(1155, 700))

	RenderToolbar(myApp, myWindow, updater, configurator)
}

func RenderToolbar(app fyne.App, mainWindow fyne.Window, updater *updater.Updater, configurator *config.Configurator) {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), getHomeContent(app, updater)),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), getSettingsContent(configurator)),
		container.NewTabItemWithIcon("News", theme.DocumentIcon(), getNewsContent()),
		container.NewTabItemWithIcon("Updates", theme.DownloadIcon(), getUpdateContent(app, updater)),
		container.NewTabItemWithIcon("About", theme.ComputerIcon(), getAboutContent()),
	)

	tabs.SetTabLocation(container.TabLocationTop)
	mainWindow.SetContent(tabs)
	mainWindow.RequestFocus()
	mainWindow.ShowAndRun()
}

