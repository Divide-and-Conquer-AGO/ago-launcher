package gui

import (
	"ago-launcher/config"
	"ago-launcher/news"
	"ago-launcher/quotes"
	"ago-launcher/updater"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	fynetooltip "github.com/dweymouth/fyne-tooltip"
)

func InitGUI(updater *updater.Updater, configurator *config.Configurator, quoter *quotes.Qouter, newsReader *news.NewsReader) {
	myApp := app.NewWithID("divide.and.conquer.ago")

	// Set the theme
	myApp.Settings().SetTheme(&AgoTheme{})

	// Create the default window
	myWindow := myApp.NewWindow("AGO Launcher")

	// Set the size and focus
	// myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(1155, 700))

	// Render the main toolbar
	RenderToolbar(myApp, myWindow, updater, configurator, quoter, newsReader)
}

func RenderToolbar(app fyne.App, mainWindow fyne.Window, updater *updater.Updater, configurator *config.Configurator, quoter *quotes.Qouter, newsReader *news.NewsReader) {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), getHomeContent(app, updater, quoter)),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), getSettingsContent(configurator)),
		container.NewTabItemWithIcon("News", theme.DocumentIcon(), getNewsContent(newsReader)),
		container.NewTabItemWithIcon("Updates", theme.DownloadIcon(), getUpdateContent(app, updater)),
		container.NewTabItemWithIcon("About", theme.ComputerIcon(), getAboutContent()),
	)

	bg := canvas.NewImageFromFile("background.png")
    bg.FillMode = canvas.ImageFillStretch // or ImageFillContain

    content := container.NewStack(
        bg,
        tabs,
    )

	tabs.SetTabLocation(container.TabLocationTop)
	mainWindow.SetContent(fynetooltip.AddWindowToolTipLayer(content, mainWindow.Canvas()))
	mainWindow.RequestFocus()
	mainWindow.ShowAndRun()
}

