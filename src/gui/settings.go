package gui

import (
	"ago-launcher/config"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func getSettingsContent(configurator *config.Configurator) fyne.CanvasObject {
	// Mod Settings (AGO.cfg)
	modSettingsTabs := container.NewAppTabs(
		container.NewTabItem("Debug", getDebugInputs(configurator)),
		container.NewTabItem("Sorting", getSortingInputs(configurator)),
		container.NewTabItem("Limits", getLimitsInputs(configurator)),
		container.NewTabItem("Saving", getSavingInputs(configurator)),
		container.NewTabItem("Info", getInfoInputs(configurator)),
		container.NewTabItem("Scripts", getScriptsInputs(configurator)),
		container.NewTabItem("Battle", getBattleInputs(configurator)),
		container.NewTabItem("Difficulty", getDifficultyInputs(configurator)),
	)
	// Game Settings (TATW.cfg)
	gameSettingsTabs := container.NewAppTabs(
		container.NewTabItem("Video", getGameInputs(configurator)),
	)
	settingsTabs := container.NewAppTabs(
		container.NewTabItem("Game Settings", gameSettingsTabs),
		container.NewTabItem("Mod Settings", modSettingsTabs),
	)
	saveButton := widget.NewButton("Save Settings", func() {
		configurator.WriteConfigToFile("AGO.cfg", configurator.AGOConfigFile, &configurator.AGOConfig)
		configurator.WriteConfigToFile("TATW.cfg", configurator.ModConfigFile, &configurator.ModConfig)
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

func getSortingInputs(configurator *config.Configurator) fyne.CanvasObject {
	option1 := widget.NewCheckWithData("Automatic stack sorting", binding.BindBool(&configurator.AGOConfig.Sorting.EnableSorting))
	option2 := widget.NewCheckWithData("Sort player stacks automatically", binding.BindBool(&configurator.AGOConfig.Sorting.SortPlayer))

	// Sort Mode Selection
	label := widget.NewLabel("Sort Algorithm Priority")
	label.TextStyle = fyne.TextStyle{Bold: true}

	sortOptions := []string{"eduType", "category", "class", "soldierCount", "experience", "categoryClass", "aiUnitValue"}

	sortMode1 := widget.NewSelect(sortOptions, func(selected string) {
		fmt.Println("Selected sort mode 1:", selected)
		for i, v := range sortOptions {
			if v == selected {
				fmt.Println("Selected index:", i+1)
				configurator.AGOConfig.Sorting.SortMode1 = i + 1
				break
			}
		}
	})
	// We use -1 because Lua indexes start from 1
	sortMode1.SetSelectedIndex(configurator.AGOConfig.Sorting.SortMode1 - 1)

	sortMode2 := widget.NewSelect(sortOptions, func(selected string) {
		fmt.Println("Selected sort mode 2:", selected)
		for i, v := range sortOptions {
			if v == selected {
				fmt.Println("Selected index:", i+1)
				configurator.AGOConfig.Sorting.SortMode2 = i + 1
				break
			}
		}
	})
	// We use -1 because Lua indexes start from 1
	sortMode2.SetSelectedIndex(configurator.AGOConfig.Sorting.SortMode2 - 1)

	sortMode3 := widget.NewSelect(sortOptions, func(selected string) {
		fmt.Println("Selected sort mode 3:", selected)
		for i, v := range sortOptions {
			if v == selected {
				fmt.Println("Selected index:", i+1)
				configurator.AGOConfig.Sorting.SortMode3 = i + 1
				break
			}
		}
	})
	// We use -1 because Lua indexes start from 1
	sortMode3.SetSelectedIndex(configurator.AGOConfig.Sorting.SortMode3 - 1)

	// Container
	content := container.NewVBox(
		option1, option2, label, sortMode1, sortMode2, sortMode3,
	)

	return content
}

func getLimitsInputs(configurator *config.Configurator) fyne.CanvasObject {
	guildCooldownSpin := MakeSpinBox(
		"Guild Cooldown",
		func() int { return configurator.AGOConfig.Limits.GuildCooldown },
		func(v int) { configurator.AGOConfig.Limits.GuildCooldown = v },
	)

	maxAncillariesSpin := MakeSpinBox(
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

func getSavingInputs(configurator *config.Configurator) fyne.CanvasObject {
	option1 := widget.NewCheckWithData("Post Battle Saving", binding.BindBool(&configurator.AGOConfig.Saving.PostBattleSaving))

	// Container
	content := container.NewVBox(
		option1,
	)

	return content
}

func getInfoInputs(configurator *config.Configurator) fyne.CanvasObject {
	option1 := widget.NewCheckWithData("Hide Army Info", binding.BindBool(&configurator.AGOConfig.Info.HideArmyInfo))
	option2 := widget.NewCheckWithData("AI Raid Notification", binding.BindBool(&configurator.AGOConfig.Info.AIRaidNotification))
	watchtowerRadius := MakeSpinBox(
		"Watchtower Radius",
		func() int { return configurator.AGOConfig.Info.WatchtowerRadius },
		func(v int) { configurator.AGOConfig.Info.WatchtowerRadius = v },
	)
	content := container.NewVBox(
		option1, option2, watchtowerRadius,
	)
	return content
}

func getScriptsInputs(configurator *config.Configurator) fyne.CanvasObject {
	naturalDisasters := widget.NewCheckWithData("Natural Disasters", binding.BindBool(&configurator.AGOConfig.Scripts.NaturalDisasters))
	randomAAAIStart := widget.NewCheckWithData("Random AA AI Start", binding.BindBool(&configurator.AGOConfig.Scripts.RandomAaAiStart))
	mergeDolAmroth := widget.NewCheckWithData("Merge Dol Amroth", binding.BindBool(&configurator.AGOConfig.Scripts.MergeDolAmroth))
	randomizedStart := widget.NewCheckWithData("Randomized Start", binding.BindBool(&configurator.AGOConfig.Scripts.RandomizedStart))
	shatteredAlliances := widget.NewCheckWithData("Shattered Alliances", binding.BindBool(&configurator.AGOConfig.Scripts.ShatteredAlliances))
	lastStandArmies := widget.NewCheckWithData("Last Stand Armies", binding.BindBool(&configurator.AGOConfig.Scripts.LastStandArmies))
	content := container.NewVBox(
		naturalDisasters, randomAAAIStart, mergeDolAmroth, randomizedStart, shatteredAlliances, lastStandArmies,
	)
	return content
}

func getBattleInputs(configurator *config.Configurator) fyne.CanvasObject {
	noDefaultSkirmish := widget.NewCheckWithData("No Default Skirmish", binding.BindBool(&configurator.AGOConfig.Battle.NoDefaultSkirmish))
	defaultBattleSpeed := MakeSpinBox(
		"Default Battle Speed",
		func() int { return configurator.AGOConfig.Battle.DefaultBattleSpeed },
		func(v int) { configurator.AGOConfig.Battle.DefaultBattleSpeed = v },
	)
	content := container.NewVBox(
		noDefaultSkirmish, defaultBattleSpeed,
	)
	return content
}

func getDifficultyInputs(configurator *config.Configurator) fyne.CanvasObject {
	aggressiveRebels := widget.NewCheckWithData("Aggressive Rebels", binding.BindBool(&configurator.AGOConfig.Difficulty.AggressiveRebels))
	aiFreeGenerals := widget.NewCheckWithData("AI Free Generals", binding.BindBool(&configurator.AGOConfig.Difficulty.AIFreeGenerals))
	content := container.NewVBox(
		aggressiveRebels, aiFreeGenerals,
	)
	return content
}

func getGameInputs(configurator *config.Configurator) fyne.CanvasObject {
	option1 := widget.NewCheckWithData("Borderless Window", binding.BindBool(&configurator.ModConfig.Video.BorderlessWindow))
	option2 := widget.NewCheckWithData("Windowed", binding.BindBool(&configurator.ModConfig.Video.Windowed))
	option3 := widget.NewCheckWithData("Bloom", binding.BindBool(&configurator.ModConfig.Video.Bloom))

	// option4 := widget.NewEntryWithData(binding.BindString(&configurator.ModConfig.Video.BattleResolution))
	// option5 := widget.NewEntryWithData(binding.BindString(&configurator.ModConfig.Video.CampaignResolution))
	option4 := MakeStringBindingField("Battle Resolution", configurator.ModConfig.Video.BattleResolution)
	option5 := MakeStringBindingField("Campaign Resolution", configurator.ModConfig.Video.CampaignResolution)

	// Container
	content := container.NewVBox(
		option1, option2, option3, option4, option5,
	)

	return content
}