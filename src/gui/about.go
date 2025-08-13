package gui

import (
	"ago-launcher/utils"
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getAboutContent() fyne.CanvasObject {
	// Logo
	logo := canvas.NewImageFromResource(resourceTolkienPng)
	logo.FillMode = canvas.ImageFillOriginal
	logoContainer := container.NewCenter(logo)

	tolkienText1 := canvas.NewText("In memory of J.R.R. Tolkien", color.White)
	tolkienText1.TextSize = 18
	tolkienText1.TextStyle = fyne.TextStyle{Italic: true}
	tolkienContainer1 := container.NewCenter(tolkienText1)

	tolkienText2 := canvas.NewText("January 3, 1892 - September 2, 1973", color.White)
	tolkienText2.TextSize = 16
	tolkienText2.TextStyle = fyne.TextStyle{Italic: true}
	tolkienContainer2 := container.NewCenter(tolkienText2)

	dummyContainer := container.NewCenter(canvas.NewText("", color.Transparent))

	// Text
	// Title
	discordURL, err := url.Parse("https://discord.gg/yVHm7kBTAY")
	if err != nil {
		utils.Logger().Println("[About]invalid discord url")
	}
	discordText := widget.NewHyperlink("Join our Discord", discordURL)
	discordText.TextStyle = fyne.TextStyle{Bold: true}
	discordContainer := container.NewCenter(discordText)

	// Website Link
	soundsOfMiddleEarth, err := url.Parse("https://sounds-of-middle-earth.com/")
	if err != nil {
		utils.Logger().Println("[About]invalid website url")
	}
	soundsOfMiddleEarthText := widget.NewHyperlink("Sounds of Middle-earth", soundsOfMiddleEarth)
	soundsOfMiddleEarthText.TextStyle = fyne.TextStyle{Bold: true}
	soundsOfMiddleEarthContainer := container.NewCenter(soundsOfMiddleEarthText)

	websiteURL, err := url.Parse("https://github.com/EddieEldridge/ago-launcher/tree/main")
	if err != nil {
		utils.Logger().Println("[About]invalid website url")
	}
	websiteText := widget.NewHyperlink("Source Code", websiteURL)
	websiteText.TextStyle = fyne.TextStyle{Bold: true}
	websiteContainer := container.NewCenter(websiteText)

	// Container
	content := container.NewVBox(
		logoContainer, tolkienContainer1, tolkienContainer2, dummyContainer, discordContainer, soundsOfMiddleEarthContainer, websiteContainer,
	)
	return content
}
