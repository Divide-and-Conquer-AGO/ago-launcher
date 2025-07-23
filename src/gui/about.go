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

	tolkienText := canvas.NewText("January 3, 1892 - September 2, 1973", color.White)
	tolkienText.TextSize = 14
	tolkienText.TextStyle = fyne.TextStyle{Italic: true}
	tolkienContainer := container.NewCenter(tolkienText)

	// Text
	// Title
	discordURL, err := url.Parse("https://discord.gg/yVHm7kBTAY")
	if err != nil {
		utils.Logger().Println("invalid discord url")
	}
	discordText := widget.NewHyperlink("Join our Discord", discordURL)
	discordText.TextStyle = fyne.TextStyle{Bold: true}
	discordContainer := container.NewCenter(discordText)

	// Website Link
	soundsOfMiddleEarth, err := url.Parse("https://sounds-of-middle-earth.com/")
	if err != nil {
		utils.Logger().Println("invalid website url")
	}
	soundsOfMiddleEarthText := widget.NewHyperlink("Sounds of Middle-earth", soundsOfMiddleEarth)
	soundsOfMiddleEarthText.TextStyle = fyne.TextStyle{Bold: true}
	soundsOfMiddleEarthContainer := container.NewCenter(soundsOfMiddleEarthText)

	titleText := canvas.NewText("Launcher created by Medik", color.White)
	titleText.TextSize = 12
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	titleContainer := container.NewCenter(titleText)

	websiteURL, err := url.Parse("https://github.com/EddieEldridge/ago-launcher/tree/main")
	if err != nil {
		utils.Logger().Println("invalid website url")
	}
	websiteText := widget.NewHyperlink("Source Code", websiteURL)
	websiteText.TextStyle = fyne.TextStyle{Bold: true}
	websiteContainer := container.NewCenter(websiteText)

	// Container
	content := container.NewVBox(
		logoContainer, tolkienContainer, discordContainer, soundsOfMiddleEarthContainer, websiteContainer, titleContainer, 
	)
	return content
}
