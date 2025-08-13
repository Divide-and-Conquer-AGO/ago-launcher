package gui

import (
	"ago-launcher/updater"
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func getUpdateContent(app fyne.App, window fyne.Window, updtr *updater.Updater) fyne.CanvasObject {
	// Table header
	headerVersion := widget.NewLabelWithStyle("Version", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	headerSavegame := widget.NewLabelWithStyle("Savegame Compatible", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	headerStatus := widget.NewLabelWithStyle("Status", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	headerDownload := widget.NewLabelWithStyle("Download", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	header := container.NewGridWithColumns(4, headerVersion, headerSavegame, headerStatus, headerDownload)

	// Table rows
	var tableRows []fyne.CanvasObject
	for _, v := range updtr.AvailableVersions.ModVersions {
		// Version label 
		versionLabel := widget.NewLabel(v.Version)
		versionLabel.Alignment = fyne.TextAlignLeading

		// Savegame compatibility label - center aligned (short Yes/No values)
		savegameLabel := widget.NewLabel("❌")
		if v.SaveGameCompatible {
			savegameLabel.SetText("✅")
		}
		savegameLabel.Alignment = fyne.TextAlignCenter

		// Status label 
		statusLabel := widget.NewLabel("")
		statusLabel.Alignment = fyne.TextAlignCenter

		// Download URL as a clickable hyperlink
		parsedUrl, err := url.Parse(v.Url)
		if err != nil {
			parsedUrl = &url.URL{}
		}
		downloadLabel := widget.NewHyperlink("Manual", parsedUrl)
		downloadLabel.Alignment = fyne.TextAlignCenter

		switch v.Version {
			case updtr.CurrentVersion.Version:
				versionLabel.TextStyle = fyne.TextStyle{Bold: true}
				statusLabel.SetText("Current")
				statusLabel.TextStyle = fyne.TextStyle{Bold: true}
			case updtr.LatestVersion.Version:
				versionLabel.TextStyle = fyne.TextStyle{Bold: true}
				statusLabel.SetText("Latest")
				statusLabel.TextStyle = fyne.TextStyle{Bold: true}
			default:
				versionLabel.TextStyle = fyne.TextStyle{Bold: true}
				statusLabel.SetText("N/A")
				statusLabel.TextStyle = fyne.TextStyle{Bold: true}
		}

		row := container.NewGridWithColumns(4, versionLabel, savegameLabel, statusLabel, downloadLabel)
		tableRows = append(tableRows, row)
	}

	// Combine header and rows into a table
	table := container.NewVBox(header)
	for _, row := range tableRows {
		table.Add(row)
	}

	// Wrap the table in a scroll container
	scrollableTable := container.NewScroll(table)
	scrollableTable.SetMinSize(fyne.NewSize(500, 300)) // Increased width for new column

	// Buttons - stacked vertically
	checkUpdateButton := widget.NewButton("Check for updates", func() {
		updtr.GetCurrentModVersion()
		updtr.CheckForUpdate()
		if updtr.UpdateAvailable {
			app.SendNotification(&fyne.Notification{
				Title:   "Update Available",
				Content: "A new mod version is available: " + updtr.LatestVersion.Version,
			})
		} else {
			app.SendNotification(&fyne.Notification{
				Title:   "No Updates",
				Content: "You are already on the latest version: " + updtr.CurrentVersion.Version,
			})
		}
	})

	var buttonBox *fyne.Container
	if updtr.UpdateAvailable {
		startUpdateButton := widget.NewButton("Update to latest version", func() {
			getUpdaterModal(updtr, window)
		})
		buttonBox = container.NewVBox(checkUpdateButton, startUpdateButton)
		app.SendNotification(&fyne.Notification{
			Title:   "Update Available",
			Content: "A new mod version is available: " + updtr.LatestVersion.Version,
		})
	} else {
		buttonBox = container.NewVBox(checkUpdateButton)
	}

	// Create the final layout with scrollable content
	content := container.NewBorder(
		nil,             // top
		buttonBox,       // bottom - buttons directly at bottom
		nil,             // left
		nil,             // right
		scrollableTable, // center - scrollable table
	)

	return content
}

func getUpdaterModal(updtr *updater.Updater, parentWindow fyne.Window) {
	// Check for compatibility issues first, before showing the update window
	warnAboutUpdates(updtr, parentWindow, func(proceed bool) {
		if !proceed {
			return // User cancelled, don't show update window at all
		}
		
		// User confirmed or no compatibility issues, show update window
		updateWindow := fyne.CurrentApp().NewWindow("Updater")
		updateWindow.Resize(fyne.NewSize(1155, 300))
		updateWindow.RequestFocus()
		updateWindow.CenterOnScreen()
		updateLabel := widget.NewLabel("Starting update process...")
		progressBar := widget.NewProgressBar()
		statusLabel := widget.NewLabel("")
		downloadProgressLabel := widget.NewLabel("")
		downloadProgressBar := widget.NewProgressBar()
		updateWindow.SetContent(container.NewVBox(
			container.NewCenter(updateLabel),
			container.NewCenter(statusLabel),
			progressBar,
			container.NewCenter(downloadProgressLabel),
			downloadProgressBar,
		))
		updateWindow.Show()

		go func() {
			// Proceed with updates
			err := updtr.ApplyUpdatesSequentially(".", func(idx, total int, v updater.ModVersion) {
				fyne.Do(func() {
					updateLabel.TextStyle = fyne.TextStyle{Bold: true}
					updateLabel.SetText(fmt.Sprintf("Applying update %d of %d: %s", idx, total, v.Version))

					progressBar.SetValue(float64(idx-1) / float64(total))

					statusLabel.TextStyle = fyne.TextStyle{Bold: true}
					statusLabel.SetText(fmt.Sprintf("Downloading %s...", v.Version))

					// Reset download progress for new file
					downloadProgressLabel.SetText("Preparing download...")
					downloadProgressBar.SetValue(0)

					updateLabel.Refresh()
					statusLabel.Refresh()
					downloadProgressLabel.Refresh()
				})
			}, func(completed, total int64, percent float64) {
				// Download and extraction progress callback
				fyne.Do(func() {
					if completed == -1 && total == 0 && percent == 0 {
						// Special signal for extraction start
						statusLabel.SetText("Download complete, starting extraction...")
						downloadProgressLabel.SetText("Preparing to extract files...")
						downloadProgressBar.SetValue(1.0) // Show download as complete
					} else if completed >= 0 && total > 0 && completed <= total {
						// Download phase - using byte progress
						if total > 1024*1024 { // If size is larger than 1MB, it's likely download
							downloadProgressLabel.SetText(fmt.Sprintf("Downloaded: %.1f MB / %.1f MB (%.1f%%)",
								float64(completed)/(1024*1024),
								float64(total)/(1024*1024),
								percent*100))
							downloadProgressBar.SetValue(percent)
							statusLabel.SetText("Downloading...")
						} else {
							// Extraction phase - using file count progress
							downloadProgressLabel.SetText(fmt.Sprintf("Extracted: %d / %d files (%.1f%%)",
								completed, total, percent*100))
							downloadProgressBar.SetValue(percent)
							statusLabel.SetText("Extracting files...")
						}
					} else if completed > 0 && total == 0 {
						// Download with unknown size
						downloadProgressLabel.SetText(fmt.Sprintf("Downloaded: %.1f MB", float64(completed)/(1024*1024)))
						downloadProgressBar.SetValue(0)
						statusLabel.SetText("Downloading...")
					} else {
						// Fallback
						downloadProgressLabel.SetText(fmt.Sprintf("Processing: %d / %d", completed, total))
						downloadProgressBar.SetValue(percent)
					}
					downloadProgressLabel.Refresh()
					downloadProgressBar.Refresh()
					statusLabel.Refresh()
				})
			})
			if err != nil {
				fyne.Do(func() {
					statusLabel.TextStyle = fyne.TextStyle{Bold: true}
					statusLabel.SetText("Update failed: " + err.Error() + " Please report this issue in the Discord, providing your AGO_Launcher.log file..")
					downloadProgressLabel.SetText("Update failed!")
					statusLabel.Refresh()
					downloadProgressLabel.Refresh()
					parentWindow.Canvas().Content().Refresh()
				})
			} else {
				fyne.Do(func() {
					progressBar.SetValue(1.0)
					statusLabel.TextStyle = fyne.TextStyle{Bold: true}
					statusLabel.SetText("All updates complete!")
					downloadProgressLabel.SetText("Download complete!")
					statusLabel.Refresh()
					downloadProgressLabel.Refresh()
					parentWindow.Canvas().Content().Refresh()
				})
			}
		}()
	})
}

func warnAboutUpdates(updtr *updater.Updater, parentWindow fyne.Window, callback func(proceed bool)) {
	// Check for save game compatibility before starting updates
	hasIncompatibleUpdates := false
	incompatibleVersions := []string{}
	
	// Get updates to check compatibility
	updates, err := updtr.GetUpdatesToApply()
	if err != nil {
		// Show error dialog if we can't check updates
		dialog.ShowError(fmt.Errorf("failed to check updates: %v", err), parentWindow)
		callback(false)
		return
	}

	// Check each update for save game compatibility
	for _, update := range updates {
		if !update.SaveGameCompatible {
			hasIncompatibleUpdates = true
			incompatibleVersions = append(incompatibleVersions, update.Version)
		}
	}

	// If no incompatible updates, proceed without warning
	if !hasIncompatibleUpdates {
		callback(true)
		return
	}

	// Show warning dialog for incompatible updates
	warningText := "⚠️ Save Game Compatibility Warning ⚠️\n\n"
	warningText += "The following updates may not be compatible with existing save games:\n\n"
	for _, version := range incompatibleVersions {
		warningText += "• " + version + "\n"
	}
	warningText += "\nContinuing will likely cause existing campaigns to become unstable or unusable.\n"
	warningText += "Do you want to continue with the update?"

	dialog.ShowConfirm("Save Game Compatibility Warning", warningText, callback, parentWindow)
}