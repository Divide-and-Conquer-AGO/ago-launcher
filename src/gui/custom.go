package gui

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	ttwidget "github.com/dweymouth/fyne-tooltip/widget"
)

func MakeSpinBox(labelText string, tooltip string, get func() int, set func(int)) fyne.CanvasObject {
	val := get()

	entry := widget.NewEntry()
	entry.SetText(fmt.Sprintf("%d", val))
	entry.OnChanged = func(s string) {
		if v, err := strconv.Atoi(s); err == nil {
			val = v
			set(v)
		}
	}

	inc := ttwidget.NewButton("+", func() {
		val++
		entry.SetText(fmt.Sprintf("%d", val))
		set(val)
	})
	label := ttwidget.NewLabel(labelText)
	dec := ttwidget.NewButton("-", func() {
		val--
		entry.SetText(fmt.Sprintf("%d", val))
		set(val)
	})
	inc.SetToolTip(tooltip)
	dec.SetToolTip(tooltip)
	label.SetToolTip(tooltip)

	spinRow := container.NewHBox(dec, entry, inc)

	content := container.NewVBox(
		label,
		spinRow,
	)

	return content
}

func MakeStringBindingField(labelText string, value string, tooltip string) fyne.CanvasObject {
	entry := widget.NewEntryWithData(binding.BindString(&value))
	label := ttwidget.NewLabel(labelText)
	label.SetToolTip(tooltip)

	content := container.NewVBox(
		label,
		entry,
	)

	return content
}

// Create a notification icon with a badge number
func NotificationIconWithBadge(icon fyne.Resource, badgeNumber int) fyne.CanvasObject {
    iconImg := widget.NewIcon(icon)

    // Only show badge if number > 0
    if badgeNumber > 0 {
        badge := canvas.NewCircle(color.NRGBA{R: 220, G: 0, B: 0, A: 255})
        badge.Resize(fyne.NewSize(18, 18))
        badge.Move(fyne.NewPos(18, 0)) // Adjust position as needed

        badgeText := canvas.NewText(strconv.Itoa(badgeNumber), color.White)
        badgeText.TextSize = 12
        badgeText.Alignment = fyne.TextAlignCenter
        badgeText.TextStyle = fyne.TextStyle{Bold: true}
        badgeText.Move(fyne.NewPos(18, 0)) // Adjust position as needed

        return container.NewStack(
            iconImg,
            badge,
            badgeText,
        )
    }
    return iconImg
}