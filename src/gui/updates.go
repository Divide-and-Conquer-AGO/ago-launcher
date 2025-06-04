package gui

import (
	"ago-launcher/updater"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getUpdateContent(app fyne.App, updater *updater.Updater) fyne.CanvasObject {
	// Check for Updates
	updateButtonLabel := "Check for updates"
	updateButton := widget.NewButton(updateButtonLabel, func() {
		newVersion, updateAvailable, err := updater.CheckForUpdate()
		if err != nil {
			fmt.Println(err)
		}
		if updateAvailable {
			app.SendNotification(fyne.NewNotification("New update available!", newVersion.Version))
		} else {
			app.SendNotification(fyne.NewNotification("You are up to date!", updater.CurrentVersion.Version))
		}
	})
	// Container
	content := container.NewVBox(
		updateButton,
	)

	return content
}
