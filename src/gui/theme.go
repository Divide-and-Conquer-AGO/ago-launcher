package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type AgoTheme struct{}

var _ fyne.Theme = (*AgoTheme)(nil)

func (m AgoTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	// Main Background
	if name == theme.ColorNameBackground {
		return color.RGBA{R: 34, G: 34, B: 34, A: 255}
	}

	// Primary Colour
	if name == theme.ColorNamePrimary {
		return color.RGBA{R: 255, G: 177, B: 68, A: 255}
	}

	// Button Background
	if name == theme.ColorNameButton {
		return color.RGBA{R: 137, G: 87, B: 0, A: 255}
	}

	// Links
	if name == theme.ColorNameHyperlink {
		return color.RGBA{R: 255, G: 177, B: 68, A: 255}
	}

	// Scrollbar
	if name == theme.ColorNameScrollBar {
		return color.RGBA{R: 255, G: 177, B: 68, A: 255}
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m AgoTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m AgoTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m AgoTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
