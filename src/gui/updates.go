package gui

import (
	"ago-launcher/updater"
	"fmt"
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
		startUpdateButton := widget.NewButton("Download Update", func() {
			getUpdaterModal(updater)
		})
		objects = append(objects, startUpdateButton)
	}

	// Container
	content := container.NewVBox(
		objects...,
	)

	return content
}

func getUpdaterModal(updtr *updater.Updater) {
	updateWindow := fyne.CurrentApp().NewWindow("Updater")
	updateWindow.Resize(fyne.NewSize(1155, 300))
	updateWindow.RequestFocus()
	updateWindow.CenterOnScreen()
	updateLabel := widget.NewLabel("Starting update process...")
	progressBar := widget.NewProgressBar()
	statusLabel := widget.NewLabel("")
	updateWindow.SetContent(container.NewVBox(
		container.NewCenter(updateLabel),
		container.NewCenter(statusLabel),
		progressBar,
	))
	updateWindow.Show()
	go func() {
	err := updtr.ApplyUpdatesSequentially(".", func(idx, total int, v updater.ModVersion) {
		fyne.Do(func() {
			updateLabel.TextStyle = fyne.TextStyle{Bold: true}
			updateLabel.SetText(fmt.Sprintf("Applying update %d of %d: %s", idx, total, v.Version))

			progressBar.SetValue(float64(idx-1) / float64(total))

			statusLabel.TextStyle = fyne.TextStyle{Bold: true}
			statusLabel.SetText(fmt.Sprintf("Downloading %s...", v.Version))

			updateLabel.Refresh()
			statusLabel.Refresh()
		})
	})
	if err != nil {
		fyne.Do(func() {
			statusLabel.TextStyle = fyne.TextStyle{Bold: true}
			statusLabel.SetText("Update failed: " + err.Error())
			statusLabel.Refresh()
		})
	} else {
		fyne.Do(func() {
			progressBar.SetValue(1.0)
			statusLabel.TextStyle = fyne.TextStyle{Bold: true}
			statusLabel.SetText("All updates complete!")
			statusLabel.Refresh()
		})
	}
}()
}
