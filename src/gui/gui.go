package gui

import (
	"ago-launcher/utils"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type AppData struct {
	Quotes utils.Quotes	
}

func InitGUI(data AppData) {
	myApp := app.NewWithID("divide.and.conquer.ago")

	myWindow := myApp.NewWindow("AGO Launcher")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(1155, 700))

	RenderToolbar(myApp, myWindow, data)
}

func RenderToolbar(app fyne.App, mainWindow fyne.Window, data AppData) {
	tabs := container.NewAppTabs(
	container.NewTabItemWithIcon("Home", theme.HomeIcon(), getHomeContent(app, data)),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), getSettingsContent()),
		container.NewTabItemWithIcon("About", theme.ComputerIcon(), getAboutContent()),
	)

	tabs.SetTabLocation(container.TabLocationTop)
	mainWindow.SetContent(tabs)
	mainWindow.RequestFocus()
	mainWindow.ShowAndRun()
}

func getHomeContent(app fyne.App, data AppData) fyne.CanvasObject {
	// Logo
	logo := canvas.NewImageFromFile("icon.png")
	logo.FillMode = canvas.ImageFillOriginal
	logoContainer := container.NewCenter(logo)

	// Text
	// Title
	titleText := canvas.NewText("Divide and Conquer: AGO V3", color.White)
	titleText.TextSize = 32
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	titleContainer := container.NewCenter(titleText)

	// Quote (Quote)
	quote, err := utils.RandomElement(data.Quotes.Quotes)
	if err != nil {
		fmt.Println("error getting quote")
	}
	quoteText := canvas.NewText(quote.Quote, color.White)
	quoteText.TextSize = 16
	quoteText.TextStyle = fyne.TextStyle{Italic: true}
	quoteContainer := container.NewCenter(quoteText)

	// Quote (Author)
	authorText := canvas.NewText(quote.Author, color.White)
	authorText.TextSize = 14
	authorText.TextStyle = fyne.TextStyle{Italic: true}
	authorContainer := container.NewCenter(authorText)

	// Buttons
	quoteButton := widget.NewButton("Refresh quote", func() {
		quote, err := utils.RandomElement(data.Quotes.Quotes)
		if err != nil {
			fmt.Println("error getting quote")
		}
		quoteText.Text = quote.Quote
		authorText.Text = quote.Author
	})
	updateButton := widget.NewButton("Check for Updates", func() {
		app.SendNotification(fyne.NewNotification("Checking for updates...", ""))
	})
	launchButton := widget.NewButton("Launch Mod", func() {
		app.SendNotification(fyne.NewNotification("Launching mod...", ""))
	})
	buttonContainer := container.NewVBox(quoteButton, updateButton, launchButton)

	// Container
	content := container.NewVBox(
        logoContainer, titleContainer, quoteContainer, authorContainer, buttonContainer,
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