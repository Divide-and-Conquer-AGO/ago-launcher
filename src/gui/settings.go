package gui

import (
	"ago-launcher/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	ttwidget "github.com/dweymouth/fyne-tooltip/widget"
)

func getSettingsContent(configurator *config.Configurator) fyne.CanvasObject {
	// Mod Settings (AGO.cfg)
	modSettingsTabs := container.NewAppTabs(
		container.NewTabItem("Video", getVideoInputs(configurator)),
		container.NewTabItem("Saving", getSavingInputs(configurator)),
		container.NewTabItem("Battle", getBattleInputs(configurator)),
		container.NewTabItem("Hotseat", getHotseatInputs(configurator)),
	)

	// Save settings
	saveButton := widget.NewButton("Save Settings", func() {
		configurator.WriteConfigToFile("TATW.cfg", &configurator.ModConfig, configurator.ModConfigFile)
		configurator.WriteConfigToFile("eopData/config/gameCfg.json", &configurator.EOPConfig.GameCfg, nil)
		configurator.WriteConfigToFile("eopData/config/battlesCfg.json", &configurator.EOPConfig.BattlesCfg, nil)
	})

	// Container
	content := container.NewVBox(
		modSettingsTabs, layout.NewSpacer(), saveButton,
	)
	return content
}

func getSavingInputs(configurator *config.Configurator) fyne.CanvasObject {
	option2 := ttwidget.NewCheckWithData("Automatic Save Backup", binding.BindBool(&configurator.EOPConfig.GameCfg.IsSaveBackupEnabled))
	option2.SetToolTip("Automatically creates multiple copies of saves in case of corruption")

	content := container.NewVBox(
		option2,
	)
	return content
}

func getBattleInputs(configurator *config.Configurator) fyne.CanvasObject {
	freeCamEnabled := ttwidget.NewCheckWithData("Freecam Integration", binding.BindBool(&configurator.EOPConfig.GameCfg.IsFreecamIntegrationEnabled))
	freeCamEnabled.SetToolTip("Automatically start and close the Freecam application when the game is launched")

	content := container.NewVBox(
		freeCamEnabled,
	)
	return content
}

func getHotseatInputs(configurator *config.Configurator) fyne.CanvasObject {
	aggressiveRebels := ttwidget.NewCheckWithData("Automatically generate hotseat/historical battles", binding.BindBool(&configurator.EOPConfig.BattlesCfg.EnableAutoGeneration))
	aggressiveRebels.SetToolTip("Enable if you want to generate a historical battle each time you start a battle")

	aiFreeGenerals := ttwidget.NewCheckWithData("Automatically generate battle result files", binding.BindBool(&configurator.EOPConfig.BattlesCfg.EnableResultsTransfer))
	aiFreeGenerals.SetToolTip("Enable if you want to generate a results file from an online battle")

	content := container.NewVBox(
		aggressiveRebels, aiFreeGenerals,
	)
	return content
}

func getVideoInputs(configurator *config.Configurator) fyne.CanvasObject {
	option1 := ttwidget.NewCheckWithData("Borderless Window", binding.BindBool(&configurator.ModConfig.Video.BorderlessWindow))
	option1.SetToolTip("Enable borderless window mode")

	option2 := ttwidget.NewCheckWithData("Windowed", binding.BindBool(&configurator.ModConfig.Video.Windowed))
	option2.SetToolTip("Enable windowed mode")

	option6 := ttwidget.NewCheckWithData("Vulkan Rendering Mode (DXVK)", binding.BindBool(&configurator.EOPConfig.GameCfg.IsDXVKEnabled))
	option6.SetToolTip("Experimental: Forces Medieval 2 to use DXVK instead of DirectX for rendering. Can massively improve performance on some hardware. \nNote: The first time you use DXVK Rendering, you may experience worse performance due to compilation of shaders.\nThe second time you launch the game, assuming the shaders have compiled, performance should be much better (even better than Vanilla DirectX Rendering)")

	option3 := ttwidget.NewCheckWithData("Battle Cutscenes", binding.BindBool(&configurator.ModConfig.Game.EventCutscenes))
	option3.SetToolTip("Enable battle cutscenes which take over camera control such as those of generals getting killed and gates being broken")

	option4 := MakeStringBindingField("Battle Resolution", &configurator.ModConfig.Video.BattleResolution, "Battle resolution (e.g. 1920 1080)")

	option5 := MakeStringBindingField("Campaign Resolution", &configurator.ModConfig.Video.CampaignResolution, "Campaign resolution (e.g. 1920x1080)")

	content := container.NewVBox(
		option1, option2, option3, option6, option4, option5,
	)
	return content
}
