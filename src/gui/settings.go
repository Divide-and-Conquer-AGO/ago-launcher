package gui

import (
	"ago-launcher/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	ttwidget "github.com/dweymouth/fyne-tooltip/widget"
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
		container.NewTabItem("AGO Settings", modSettingsTabs),
	)
	saveButton := widget.NewButton("Save Settings", func() {
		configurator.WriteConfigToFile("AGO.cfg", configurator.AGOConfigFile, &configurator.AGOConfig)
		configurator.WriteConfigToFile("TATW.cfg", configurator.ModConfigFile, &configurator.ModConfig)
	})

	// Container
	content := container.NewVBox(
		settingsTabs, saveButton,
	)
	return content
}
func getDebugInputs(configurator *config.Configurator) fyne.CanvasObject {
	option1 := ttwidget.NewCheckWithData("Enable Logging", binding.BindBool(&configurator.AGOConfig.Debug.EnableLogging))
	option1.SetToolTip("Enable logging")

	option2 := ttwidget.NewCheckWithData("Developer Debug", binding.BindBool(&configurator.AGOConfig.Debug.DevDebug))
	option2.SetToolTip("Disables Fog of War, enables Perfect Spy and shows all army information")

	option3 := ttwidget.NewCheckWithData("Log to Console", binding.BindBool(&configurator.AGOConfig.Debug.LogToConsole))
	option3.SetToolTip("Direct log statements to the EOP console")

	content := container.NewVBox(
		option1, option2, option3,
	)
	return content
}

func getSortingInputs(configurator *config.Configurator) fyne.CanvasObject {
	option1 := ttwidget.NewCheckWithData("Automatic stack sorting", binding.BindBool(&configurator.AGOConfig.Sorting.EnableSorting))
	option1.SetToolTip("Enables automatic sorting of AI stacks")

	option2 := ttwidget.NewCheckWithData("Sort player stacks automatically", binding.BindBool(&configurator.AGOConfig.Sorting.SortPlayer))
	option2.SetToolTip("Automatically sort the players stacks")

	label := widget.NewLabel("Sort Algorithm Priority")
	label.TextStyle = fyne.TextStyle{Bold: true}

	sortOptions := []string{"eduType", "category", "class", "soldierCount", "experience", "categoryClass", "aiUnitValue"}

	sortMode1 := widget.NewSelect(sortOptions, func(selected string) {
		for i, v := range sortOptions {
			if v == selected {
				configurator.AGOConfig.Sorting.SortMode1 = i + 1
				break
			}
		}
	})
	sortMode1.SetSelectedIndex(configurator.AGOConfig.Sorting.SortMode1 - 1)

	sortMode2 := widget.NewSelect(sortOptions, func(selected string) {
		for i, v := range sortOptions {
			if v == selected {
				configurator.AGOConfig.Sorting.SortMode2 = i + 1
				break
			}
		}
	})
	sortMode2.SetSelectedIndex(configurator.AGOConfig.Sorting.SortMode2 - 1)

	sortMode3 := widget.NewSelect(sortOptions, func(selected string) {
		for i, v := range sortOptions {
			if v == selected {
				configurator.AGOConfig.Sorting.SortMode3 = i + 1
				break
			}
		}
	})
	sortMode3.SetSelectedIndex(configurator.AGOConfig.Sorting.SortMode3 - 1)

	content := container.NewVBox(
		option1, option2, label, sortMode1, sortMode2, sortMode3,
	)
	return content
}

func getLimitsInputs(configurator *config.Configurator) fyne.CanvasObject {
	guildCooldownSpin := MakeSpinBox(
		"Guild Cooldown",
		"Number of turns between guild offers",
		func() int { return configurator.AGOConfig.Limits.GuildCooldown },
		func(v int) { configurator.AGOConfig.Limits.GuildCooldown = v },
	)

	maxAncillariesSpin := MakeSpinBox(
		"Maximum Ancillaries",
		"Maximum number of ancillaries any character can hold",
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
	option1 := ttwidget.NewCheckWithData("Post Battle Saving", binding.BindBool(&configurator.AGOConfig.Saving.PostBattleSaving))
	option1.SetToolTip("Automatically creates a save after a battle")

	content := container.NewVBox(
		option1,
	)
	return content
}

func getInfoInputs(configurator *config.Configurator) fyne.CanvasObject {
	option1 := ttwidget.NewCheckWithData("Hide Army Info", binding.BindBool(&configurator.AGOConfig.Info.HideArmyInfo))
	option1.SetToolTip("Hides army information like banner size (Triggers a reset of the display if changed in-game, your screen may flash white after toggling this setting. Please wait.)")

	option2 := ttwidget.NewCheckWithData("AI Raid Notification", binding.BindBool(&configurator.AGOConfig.Info.AIRaidNotification))
	option2.SetToolTip("Should the player be notified whenever the AI performs an active raid on their lands")

	watchtowerRadius := MakeSpinBox(
		"Watchtower Radius",
		"Radius of watchtowers in tiles",
		func() int { return configurator.AGOConfig.Info.WatchtowerRadius },
		func(v int) { configurator.AGOConfig.Info.WatchtowerRadius = v },
	)

	content := container.NewVBox(
		option1, option2, watchtowerRadius,
	)
	return content
}

func getScriptsInputs(configurator *config.Configurator) fyne.CanvasObject {
	naturalDisasters := ttwidget.NewCheckWithData("Natural Disasters", binding.BindBool(&configurator.AGOConfig.Scripts.NaturalDisasters))
	naturalDisasters.SetToolTip("Should natural disasters such as earthquakes, forest fires, tidal waves and more randomly occur during the campaign.")

	randomAAAIStart := ttwidget.NewCheckWithData("Random AA AI Start", binding.BindBool(&configurator.AGOConfig.Scripts.RandomAaAiStart))
	randomAAAIStart.SetToolTip("AI Ar-Adunaim start at a random coastal location. Has no effect if the campaign has already been started and thus needs to be set manually in the launcher or at the main menu.")

	mergeDolAmroth := ttwidget.NewCheckWithData("Merge Dol Amroth", binding.BindBool(&configurator.AGOConfig.Scripts.MergeDolAmroth))
	mergeDolAmroth.SetToolTip("Automatically merge Dol Amroth into Gondor at game start. Has no effect if the campaign has already been started and thus needs to be set manually in the launcher or at the main menu.")

	randomizedStart := ttwidget.NewCheckWithData("Randomized Start", binding.BindBool(&configurator.AGOConfig.Scripts.RandomizedStart))
	randomizedStart.SetToolTip("Randomize the starting positions of the factions. Has no effect if the campaign has already been started and thus needs to be set manually in the launcher or at the main menu.")

	shatteredAlliances := ttwidget.NewCheckWithData("Shattered Alliances", binding.BindBool(&configurator.AGOConfig.Scripts.ShatteredAlliances))
	shatteredAlliances.SetToolTip("All factions start neutral towards each other and automatic expansion at turn 1 is disabled. Has no effect if the campaign has already been started and thus needs to be set manually in the launcher or at the main menu.")

	lastStandArmies := ttwidget.NewCheckWithData("Last Stand Armies", binding.BindBool(&configurator.AGOConfig.Scripts.LastStandArmies))
	lastStandArmies.SetToolTip("Should factions recieve a last stand army when they are close to being defeated.")

	content := container.NewVBox(
		naturalDisasters, randomAAAIStart, mergeDolAmroth, randomizedStart, shatteredAlliances, lastStandArmies,
	)
	return content
}

func getBattleInputs(configurator *config.Configurator) fyne.CanvasObject {
	noDefaultSkirmish := ttwidget.NewCheckWithData("No Default Skirmish", binding.BindBool(&configurator.AGOConfig.Battle.NoDefaultSkirmish))
	noDefaultSkirmish.SetToolTip("Disable skirmish mode for player units by default")

	defaultBattleSpeed := MakeSpinBox(
		"Default Battle Speed",
		"Default speed set at the start of a battle",
		func() int { return configurator.AGOConfig.Battle.DefaultBattleSpeed },
		func(v int) { configurator.AGOConfig.Battle.DefaultBattleSpeed = v },
	)

	content := container.NewVBox(
		noDefaultSkirmish, defaultBattleSpeed,
	)
	return content
}

func getDifficultyInputs(configurator *config.Configurator) fyne.CanvasObject {
	aggressiveRebels := ttwidget.NewCheckWithData("Aggressive Rebels", binding.BindBool(&configurator.AGOConfig.Difficulty.AggressiveRebels))
	aggressiveRebels.SetToolTip("Rebels will be more aggressive as well as attack settlements")

	aiFreeGenerals := ttwidget.NewCheckWithData("AI Free Generals", binding.BindBool(&configurator.AGOConfig.Difficulty.AIFreeGenerals))
	aiFreeGenerals.SetToolTip("AI gets free generals on large captain armies (balanced with a settlement/general ratio)")

	content := container.NewVBox(
		aggressiveRebels, aiFreeGenerals,
	)
	return content
}

func getGameInputs(configurator *config.Configurator) fyne.CanvasObject {
	option1 := ttwidget.NewCheckWithData("Borderless Window", binding.BindBool(&configurator.ModConfig.Video.BorderlessWindow))
	option1.SetToolTip("Enable borderless window mode")

	option2 := ttwidget.NewCheckWithData("Windowed", binding.BindBool(&configurator.ModConfig.Video.Windowed))
	option2.SetToolTip("Enable windowed mode")

	option3 := ttwidget.NewCheckWithData("Bloom", binding.BindBool(&configurator.ModConfig.Video.Bloom))
	option3.SetToolTip("Enable bloom effect")

	option4 := MakeStringBindingField("Battle Resolution", configurator.ModConfig.Video.BattleResolution, "Battle resolution (e.g. 1920 1080)")

	option5 := MakeStringBindingField("Campaign Resolution", configurator.ModConfig.Video.CampaignResolution, "Campaign resolution (e.g. 1920x1080)")

	content := container.NewVBox(
		option1, option2, option3, option4, option5,
	)
	return content
}
