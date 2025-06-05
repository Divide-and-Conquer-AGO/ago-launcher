package gui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func MakeSpinBox(labelText string, get func() int, set func(int)) fyne.CanvasObject {
	val := get()

	entry := widget.NewEntry()
	entry.SetText(fmt.Sprintf("%d", val))
	entry.OnChanged = func(s string) {
		if v, err := strconv.Atoi(s); err == nil {
			val = v
			set(v)
		}
	}

	inc := widget.NewButton("+", func() {
		val++
		entry.SetText(fmt.Sprintf("%d", val))
		set(val)
	})
	dec := widget.NewButton("-", func() {
		val--
		entry.SetText(fmt.Sprintf("%d", val))
		set(val)
	})

	spinRow := container.NewHBox(dec, entry, inc)

	content := container.NewVBox(
		widget.NewLabel(labelText),
		spinRow,
	)

	return content
}

func MakeStringBindingField(labelText string, value string) fyne.CanvasObject {
	entry := widget.NewEntryWithData(binding.BindString(&value))

	content := container.NewVBox(
		widget.NewLabel(labelText),
		entry,
	)

	return content
}