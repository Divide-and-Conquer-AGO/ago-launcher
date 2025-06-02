package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func InitGUI() {
	myApp := app.NewWithID("divide.and.conquer.ago")

	myWindow := myApp.NewWindow("AGO Launcher")
	myWindow.CenterOnScreen()
	myWindow.FixedSize()
	myWindow.Resize(fyne.NewSize(1155, 700))

	RenderToolbar(myApp, myWindow)
}

func RenderToolbar(app fyne.App, mainWindow fyne.Window) {
	tabs := container.NewAppTabs(
	container.NewTabItemWithIcon("Home", theme.HomeIcon(), getHomeContent(app)),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), getSettingsContent()),
		container.NewTabItemWithIcon("About", theme.ComputerIcon(), getAboutContent()),
	)

	tabs.SetTabLocation(container.TabLocationTop)
	mainWindow.SetContent(tabs)
	mainWindow.RequestFocus()
	mainWindow.ShowAndRun()
}

func getHomeContent(app fyne.App) fyne.CanvasObject {
	// Logo
	logo := canvas.NewImageFromFile("icon.png")
	logo.FillMode = canvas.ImageFillOriginal
	logoContainer := container.NewCenter(logo)

	// Text
	titleText := canvas.NewText("Divide and Conquer: AGO V3", color.White)
	titleContainer := container.NewCenter(titleText)


	// Buttons
	button := widget.NewButton("Check for Updates", func() {
		app.SendNotification(fyne.NewNotification("Checking for updates...", ""))
	})
	button2 := widget.NewButton("Launch Mod", func() {
		app.SendNotification(fyne.NewNotification("Launching mod...", ""))
	})
	buttonContainer := container.NewVBox(button, button2)

	// Container
	content := container.NewVBox(
        logoContainer, titleContainer, buttonContainer,
    )
    return content
}
func getSettingsContent() fyne.CanvasObject {
	return widget.NewLabel("Settings")
}
func getAboutContent() fyne.CanvasObject {
	img := canvas.NewImageFromFile("icon.png")
	img.FillMode = canvas.ImageFillOriginal
	text := canvas.NewText("Overlay", color.Black)
	content := container.New(layout.NewCenterLayout(), img, text)
	return content
}