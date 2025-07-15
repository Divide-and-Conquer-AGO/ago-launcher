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
	logo := canvas.NewImageFromFile("medik.png")
	logo.FillMode = canvas.ImageFillOriginal
	logoContainer := container.NewCenter(logo)

	// Text
	// Title
	titleText := canvas.NewText("Created by Medik", color.White)
	titleText.TextSize = 16
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	titleContainer := container.NewCenter(titleText)

	// Website Link
	websiteURL, err := url.Parse("https://github.com/EddieEldridge/ago-launcher/tree/main")
	if err != nil {
		utils.Logger().Println("invalid website url")
	}
	websiteText := widget.NewHyperlink("Source Code", websiteURL)
	websiteText.TextStyle = fyne.TextStyle{Bold: true}
	websiteContainer := container.NewCenter(websiteText)

	// Container
	content := container.NewVBox(
		logoContainer, titleContainer, websiteContainer,
	)
	return content
}
