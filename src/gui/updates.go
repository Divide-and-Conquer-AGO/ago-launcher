package gui

import (
	"ago-launcher/updater"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getUpdateContent(app fyne.App, window fyne.Window, updater *updater.Updater) fyne.CanvasObject {
	var objects []fyne.CanvasObject

	// Mod Version
	versionBanner := canvas.NewText("Mod Version", color.White)
	versionBanner.TextSize = 34
	versionBanner.TextStyle = fyne.TextStyle{Bold: true}
	bannerContainer := container.NewCenter(versionBanner)
	objects = append(objects, bannerContainer)

	currentVersionText := canvas.NewText("Current Version: "+updater.CurrentVersion.Version, color.White)
	currentVersionText.TextSize = 24
	currentVersionText.TextStyle = fyne.TextStyle{Bold: true}
	currVersionContainer := container.NewCenter(currentVersionText)
	objects = append(objects, currVersionContainer)

	latestVersionText := canvas.NewText("Latest Version: "+updater.LatestVersion.Version, color.White)
	latestVersionText.TextSize = 24
	latestVersionText.TextStyle = fyne.TextStyle{Bold: true}

	if updater.UpdateAvailable {
		latestVersionText.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	}

	latestVersionContainer := container.NewCenter(latestVersionText)
	objects = append(objects, latestVersionContainer)

	checkUpdateButton := widget.NewButton("Check for updates", func() {
		updater.CheckForUpdate()
		latestVersionText.Text = "Latest Version: " + updater.LatestVersion.Version
		latestVersionContainer.Refresh()
	})
	objects = append(objects, checkUpdateButton)
	if updater.UpdateAvailable {
		startUpdateButton := widget.NewButton("Start Update", func() {
			// update logic here
		})
		objects = append(objects, startUpdateButton)
	}

	// Container
	content := container.NewVBox(
		objects...,
	)

	return content
}