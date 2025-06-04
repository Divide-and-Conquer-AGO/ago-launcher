package gui

import (
	"ago-launcher/config"
	"ago-launcher/quotes"
	"ago-launcher/updater"
	"fmt"
	"image/color"
	"strconv"

	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func InitGUI(updater *updater.Updater, configurator *config.Configurator) {
	myApp := app.NewWithID("divide.and.conquer.ago")

	myWindow := myApp.NewWindow("AGO Launcher")
	// myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(1155, 700))

	RenderToolbar(myApp, myWindow, updater, configurator)
}

func RenderToolbar(app fyne.App, mainWindow fyne.Window, updater *updater.Updater, configurator *config.Configurator) {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), getHomeContent(app, updater)),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), getSettingsContent(configurator)),
		container.NewTabItemWithIcon("News", theme.DocumentIcon(), getNewsContent()),
		container.NewTabItemWithIcon("Updates", theme.DownloadIcon(), getUpdateContent(app, updater)),
		container.NewTabItemWithIcon("About", theme.ComputerIcon(), getAboutContent()),
	)

	tabs.SetTabLocation(container.TabLocationTop)
	mainWindow.SetContent(tabs)
	mainWindow.RequestFocus()
	mainWindow.ShowAndRun()
}

func getHomeContent(app fyne.App, updater *updater.Updater) fyne.CanvasObject {
	var quoter = &quotes.Qouter{}

	// Logo
	logo := canvas.NewImageFromFile("icon.png")
	logo.FillMode = canvas.ImageFillOriginal
	logoContainer := container.NewCenter(logo)

	// Text
	// Title
	titleText := canvas.NewText("Divide and Conquer: AGO V3", color.White)
	titleText.TextSize = 32
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	titleContainer := container.NewCenter(titleText)

	// Quote (Quote)
	quote, err := quoter.GetRandomQuote()
	if err != nil {
		fmt.Println("error random getting quote")
	}
	quoteText := canvas.NewText(quote.Quote, color.White)
	quoteText.TextSize = 16
	quoteText.TextStyle = fyne.TextStyle{Italic: true}
	quoteContainer := container.NewCenter(quoteText)

	// Quote (Author)
	authorText := canvas.NewText(quote.Author, color.White)
	authorText.TextSize = 14
	authorText.TextStyle = fyne.TextStyle{Italic: true}
	authorContainer := container.NewCenter(authorText)

	// Mod Version
	versionText := canvas.NewText(updater.CurrentVersion.Version, color.White)
	versionText.TextSize = 12
	versionText.TextStyle = fyne.TextStyle{Bold: true}
	versionContainer := container.NewCenter(versionText)

	// Website Link
	websiteURL, err := url.Parse("https://www.divide-and-conquer-ago.com/")
	if err != nil {
		fmt.Println("invalid website url")
	}
	websiteText := widget.NewHyperlink("divide-and-conquer-ago.com", websiteURL)
	websiteText.TextStyle = fyne.TextStyle{Bold: true}
	websiteContainer := container.NewCenter(websiteText)

	// Buttons

	// Quote Refresh
	// quoteButton := widget.NewButton("Refresh quote", func() {
	// 	quote, err := quoter.GetRandomQuote()
	// 	if err != nil {
	// 		fmt.Println("error getting random quote")
	// 	}
	// 	quoteText.Text = quote.Quote
	// 	authorText.Text = quote.Author
	// })

	// Launch Mod
	launchButton := widget.NewButton("Launch Mod", func() {
		app.SendNotification(fyne.NewNotification("Launching mod...", ""))
	})
	buttonContainer := container.NewVBox(launchButton)

	// Container
	content := container.NewVBox(
		logoContainer, titleContainer, quoteContainer, authorContainer, versionContainer, websiteContainer, buttonContainer,
	)
	return content
}

func getSettingsContent(configurator *config.Configurator) fyne.CanvasObject {
	// Mod Settings (AGO.cfg)
	modSettingsTabs := container.NewAppTabs(
		container.NewTabItem("Debug", getDebugInputs(configurator)),
		container.NewTabItem("Sorting", getSortingInputs(configurator)),
		container.NewTabItem("Limits", getLimitsInputs(configurator)),
		container.NewTabItem("Saving", widget.NewLabel("Saving")),
		container.NewTabItem("Info", widget.NewLabel("Info")),
		container.NewTabItem("Scripts", widget.NewLabel("Scripts")),
		container.NewTabItem("Battle", widget.NewLabel("Battle")),
		container.NewTabItem("Difficulty", widget.NewLabel("Difficulty")),
	)
	// Game Settings (TATW.cfg)
	gameSettingsTabs := container.NewAppTabs(
		container.NewTabItem("Debug", getDebugInputs(configurator)),
		container.NewTabItem("Sorting", getSortingInputs(configurator)),
		container.NewTabItem("Limits", getLimitsInputs(configurator)),
		container.NewTabItem("Saving", widget.NewLabel("Saving")),
		container.NewTabItem("Info", widget.NewLabel("Info")),
		container.NewTabItem("Scripts", widget.NewLabel("Scripts")),
		container.NewTabItem("Battle", widget.NewLabel("Battle")),
		container.NewTabItem("Difficulty", widget.NewLabel("Difficulty")),
	)
	settingsTabs := container.NewAppTabs(
		container.NewTabItem("Game Settings", gameSettingsTabs),
		container.NewTabItem("Mod Settings", modSettingsTabs),
	)
	saveButton := widget.NewButton("Save Settings", func() {
		configurator.WriteConfigToFile("AGO.cfg", configurator.AGOConfigFile, &configurator.AGOConfig)
		configurator.WriteConfigToFile("TAWT.cfg", configurator.ModConfigFile, &configurator.ModConfig)
		fmt.Println("Saved settings")
	})

	// Container
	content := container.NewVBox(
		settingsTabs, saveButton,
	)
	return content
}

func getDebugInputs(configurator *config.Configurator) fyne.CanvasObject {
	option1 := widget.NewCheckWithData("Enable Logging", binding.BindBool(&configurator.AGOConfig.Debug.EnableLogging))
	option2 := widget.NewCheckWithData("Developer Debug", binding.BindBool(&configurator.AGOConfig.Debug.DevDebug))
	option3 := widget.NewCheckWithData("Log to Console", binding.BindBool(&configurator.AGOConfig.Debug.LogToConsole))
	
	// Container
	content := container.NewVBox(
		option1, option2, option3,
	)

	return content
}

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

func getNewsContent() fyne.CanvasObject {
	// Check for Updates
	newsWidget := widget.NewLabel("Some news item")
	// Container
	content := container.NewVBox(
		newsWidget,
	)

	return content
}

func getSortingInputs(configurator *config.Configurator) fyne.CanvasObject {
	option1 := widget.NewCheckWithData("Automatic stack sorting", binding.BindBool(&configurator.AGOConfig.Sorting.EnableSorting))
	option2 := widget.NewCheckWithData("Sort player stacks automatically", binding.BindBool(&configurator.AGOConfig.Sorting.SortPlayer))
	
	// Sort Mode Selection
	label := widget.NewLabel("Sort Algorithm Priority")
	label.TextStyle = fyne.TextStyle{Bold: true}

	sortOptions := []string{"eduType", "category", "class", "soldierCount", "experience", "categoryClass", "aiUnitValue"}

	sortMode1 := widget.NewSelect(sortOptions,  func(selected string) {
            fmt.Println("Selected sort mode 1:", selected)
			for i, v := range sortOptions {
				if v == selected {
					fmt.Println("Selected index:", i+1)
					configurator.AGOConfig.Sorting.SortMode1 = i+1 
					break
				}
			}
        },)
	// We use -1 because Lua indexes start from 1 
	sortMode1.SetSelectedIndex(configurator.AGOConfig.Sorting.SortMode1-1)

	sortMode2 := widget.NewSelect(sortOptions,  func(selected string) {
            fmt.Println("Selected sort mode 2:", selected)
			for i, v := range sortOptions {
				if v == selected {
					fmt.Println("Selected index:", i+1)
					configurator.AGOConfig.Sorting.SortMode2 = i+1 
					break
				}
			}
        },)
	// We use -1 because Lua indexes start from 1 
	sortMode2.SetSelectedIndex(configurator.AGOConfig.Sorting.SortMode2-1)

	sortMode3 := widget.NewSelect(sortOptions,  func(selected string) {
            fmt.Println("Selected sort mode 3:", selected)
			for i, v := range sortOptions {
				if v == selected {
					fmt.Println("Selected index:", i+1)
					configurator.AGOConfig.Sorting.SortMode3 = i+1
					break
				}
			}
        },)
	// We use -1 because Lua indexes start from 1 
	sortMode3.SetSelectedIndex(configurator.AGOConfig.Sorting.SortMode3-1)
	
	// Container
	content := container.NewVBox(
		option1, option2, label, sortMode1, sortMode2, sortMode3,
	)

	return content
}

func getLimitsInputs(configurator *config.Configurator) fyne.CanvasObject {
     guildCooldownSpin := makeSpinBox(
        "Guild Cooldown",
        func() int { return configurator.AGOConfig.Limits.GuildCooldown },
        func(v int) { configurator.AGOConfig.Limits.GuildCooldown = v },
    )

    maxAncillariesSpin := makeSpinBox(
        "Maximum Ancillaries",
        func() int { return configurator.AGOConfig.Limits.MaximumAncillaries },
        func(v int) { configurator.AGOConfig.Limits.MaximumAncillaries = v },
    )

    content := container.NewVBox(
        guildCooldownSpin,
        maxAncillariesSpin,
    )

    return content
}

func getAboutContent() fyne.CanvasObject {
	img := canvas.NewImageFromFile("icon.png")
	img.FillMode = canvas.ImageFillOriginal
	text := canvas.NewText("Overlay", color.Black)
	content := container.New(layout.NewCenterLayout(), img, text)
	return content
}

func makeSpinBox(labelText string, get func() int, set func(int)) fyne.CanvasObject {
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