package gui

import (
	"ago-launcher/utils"
	"os"

	"fyne.io/fyne/v2"
)

var ResourceFaviconIco = resourceFaviconIco

// loadResource prefers a file on disk so local asset changes are visible
// without rebundling, while still falling back to embedded resources.
func loadResource(fileName string, fallback fyne.Resource) fyne.Resource {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return fallback
	}

	utils.Logger().Printf("[GUI] Loaded resource from disk: %s", fileName)
	return fyne.NewStaticResource(fileName, data)
}
