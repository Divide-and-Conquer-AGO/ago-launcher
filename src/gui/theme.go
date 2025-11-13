package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// AgoTheme defines a unified dark theme that ignores system light/dark settings.
type AgoTheme struct{}

var _ fyne.Theme = (*AgoTheme)(nil)

// ---- Color Palette ----
var (
	colorBackground = color.RGBA{R: 34, G: 34, B: 34, A: 255}   // Dark gray background
	colorPrimary    = color.RGBA{R: 191, G: 111, B: 0, A: 255}  // Amber accent
	colorButton     = color.RGBA{R: 138, G: 80, B: 0, A: 255}   // Dark amber button
	colorText       = color.RGBA{R: 255, G: 255, B: 255, A: 255} // White text
)

// ---- Color Definitions ----
func (m AgoTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	// Always use dark variant for a unified appearance
	switch name {
	case theme.ColorNameBackground:
		return colorBackground
	case theme.ColorNamePrimary:
		return colorPrimary
	case theme.ColorNameButton:
		return colorButton
	case theme.ColorNameHyperlink:
		return colorPrimary
	case theme.ColorNameScrollBar:
		return colorPrimary
	case theme.ColorNameForeground:
		return colorText
	default:
		// Fallback to dark defaults for consistency
		return theme.DefaultTheme().Color(name, theme.VariantDark)
	}
}

// ---- Icon Definitions ----
func (m AgoTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// ---- Font Definitions ----
// Make sure your fonts are bundled with `fyne bundle`
func (m AgoTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace {
		return resourceLTMuseumBlackTtf
	}
	if style.Bold {
		if style.Italic {
			return resourceLTMuseumItalicTtf
		}
		return resourceLTMuseumBoldTtf
	}
	if style.Italic {
		return resourceLTMuseumItalicTtf
	}
	return resourceLTMuseumBlackTtf
}

// ---- Size Adjustments ----
func (m AgoTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameCaptionText {
		return 16
	}
	return theme.DefaultTheme().Size(name)
}
