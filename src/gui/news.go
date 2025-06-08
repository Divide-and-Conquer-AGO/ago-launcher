package gui

import (
	"ago-launcher/news"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getNewsContent(newsReader *news.NewsReader) fyne.CanvasObject {
	// Return the accordion of new news content
	accordion := widget.NewAccordion()
	for _, item := range newsReader.NewsItems {
		if item.Published {
			title := item.Title + " - " + item.Date
			content := widget.NewRichTextFromMarkdown(item.MarkdownText)
			content.Wrapping = fyne.TextWrapWord
			scroll := container.NewVScroll(content)
			// scroll.SetMinSize(fyne.NewSize(0, 200))
			accordion.Append(widget.NewAccordionItem(title, scroll))
		}
	}

	// Container
	content := container.NewVBox(
		accordion,
	)
	

	return content
}
