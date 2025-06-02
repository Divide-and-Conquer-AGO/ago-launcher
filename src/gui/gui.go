package gui

import (
	"ago-launcher/quotes"
	"ago-launcher/updater"
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

func InitGUI() {
	myApp := app.NewWithID("divide.and.conquer.ago")

	myWindow := myApp.NewWindow("AGO Launcher")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
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
	var updater = &updater.Updater{}
	var quoter = &quotes.Qouter{}

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
	quote, err := quoter.GetRandomQuote()
	if err != nil {
		fmt.Println("error random getting quote")
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

	// Mod Version
	modVersion, err := updater.GetModVersion()
	if err != nil {
		fmt.Println(err)
	}
	versionText := canvas.NewText(modVersion.Version, color.White)
	versionText.TextSize = 12
	versionText.TextStyle = fyne.TextStyle{Bold: true}
	versionContainer := container.NewCenter(versionText)

	// Buttons
	quoteButton := widget.NewButton("Refresh quote", func() {
		quote, err := quoter.GetRandomQuote()
		if err != nil {
			fmt.Println("error getting random quote")
		}
		quoteText.Text = quote.Quote
		authorText.Text = quote.Author
	})
	updateButton := widget.NewButton("Check for Updates", func() {
		app.SendNotification(fyne.NewNotification("Checking for updates...", ""))
		updater.CheckForUpdate()
	})
	launchButton := widget.NewButton("Launch Mod", func() {
		app.SendNotification(fyne.NewNotification("Launching mod...", ""))
	})
	buttonContainer := container.NewVBox(quoteButton, updateButton, launchButton)

	// Container
	content := container.NewVBox(
        logoContainer, titleContainer, quoteContainer, authorContainer, versionContainer, buttonContainer,
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