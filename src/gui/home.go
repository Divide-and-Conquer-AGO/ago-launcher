package gui

import (
	"ago-launcher/quotes"
	"ago-launcher/updater"
	"ago-launcher/utils"
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func getHomeContent(app fyne.App, updater *updater.Updater, quoter *quotes.Qouter) fyne.CanvasObject {
	// Logo
	logo := canvas.NewImageFromResource(resourceIconPng)
	logo.FillMode = canvas.ImageFillOriginal
	logoContainer := container.NewCenter(logo)

	// Text
	// Title
	titleText := canvas.NewText("Divide and Conquer: AGO", color.White)
	titleText.TextSize = 32
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	titleContainer := container.NewCenter(titleText)

	// Quote (Quote)
	   quote, err := quoter.GetRandomQuote()
	   if err != nil {
			   utils.Logger().Println("error random getting quote")
	   }
	quoteText := canvas.NewText(quote.Quote, color.White)
	quoteText.TextSize = 18
	quoteText.TextStyle = fyne.TextStyle{Italic: true}
	quoteContainer := container.NewCenter(quoteText)

	// Quote (Author)
	authorText := canvas.NewText(quote.Author, color.White)
	authorText.TextSize = 16
	authorText.TextStyle = fyne.TextStyle{Italic: true}
	authorContainer := container.NewCenter(authorText)

	// Mod Version
	versionText := canvas.NewText("Version: " + updater.CurrentVersion.Version, color.White)
	versionText.TextSize = 14
	versionText.TextStyle = fyne.TextStyle{Bold: true}
	versionContainer := container.NewCenter(versionText)

	// Website Link
	websiteURL, err := url.Parse("https://www.divide-and-conquer-ago.com/")
	if err != nil {
			utils.Logger().Println("invalid website url")
	}
	websiteText := widget.NewHyperlink("www.divide-and-conquer-ago.com", websiteURL)
	websiteText.TextStyle = fyne.TextStyle{Bold: true}
	websiteContainer := container.NewCenter(websiteText)

	// Buttons

	// Quote Refresh
	// quoteButton := widget.NewButton("Refresh quote", func() {
	// 	quote, err := quoter.GetRandomQuote()
	// 	if err != nil {
	// 		utils.Logger().Println("error getting random quote")
	// 	}
	// 	quoteText.Text = quote.Quote
	// 	authorText.Text = quote.Author
	// })

	// Launch Mod
	launchButton := widget.NewButton("Launch Mod", func() {
		utils.RunExecutable("M2TWEOP_GUI.exe")
	})
	buttonContainer := container.NewVBox(launchButton)

	// Container
	content := container.NewVBox(
		logoContainer, titleContainer, quoteContainer, authorContainer, versionContainer, websiteContainer, layout.NewSpacer(), buttonContainer,
	)
	return content
}
