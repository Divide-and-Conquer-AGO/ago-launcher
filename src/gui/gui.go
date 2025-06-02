package gui

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func InitGUI() {
	myApp := app.New()

	myWindow := myApp.NewWindow("AGO Launcher")
	myWindow.CenterOnScreen()
	myWindow.FixedSize()
	myWindow.Resize(fyne.NewSize(1500, 900))

	RenderToolbar(myApp, myWindow)
}

func RenderToolbar(app fyne.App, mainWindow fyne.Window) {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), GetHomeContent()),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), GetSettingsContent()),
		container.NewTabItemWithIcon("About", theme.ComputerIcon(), GetAboutContent()),
	)

	tabs.SetTabLocation(container.TabLocationTop)
	mainWindow.SetContent(tabs)
	mainWindow.RequestFocus()
	mainWindow.ShowAndRun()
}

func GetHomeContent() fyne.CanvasObject {
	return widget.NewLabel("Home")
}
func GetSettingsContent() fyne.CanvasObject {
	return widget.NewLabel("Settings")
}
func GetAboutContent() fyne.CanvasObject {
	img := canvas.NewImageFromFile("icon.png")
	img.FillMode = canvas.ImageFillOriginal
	text := canvas.NewText("Overlay", color.Black)
	content := container.New(layout.NewCenterLayout(), img, text)
	return content
}