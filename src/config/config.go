package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type Configurator struct {
	AGOConfigFile *ini.File
	ModConfigFile *ini.File

	AGOConfig    AGOConfig
	ModConfig    ModConfig
}

type ModConfig struct {
	Video struct {
		BorderlessWindow   bool   `ini:"borderless_window"`
		Windowed           bool   `ini:"windowed"`
		BattleResolution   string `ini:"battle_resolution"`
		CampaignResolution string `ini:"campaign_resolution"`
		Bloom              bool   `ini:"bloom"`
	} `ini:"video"`
}

type AGOConfig struct {
	Debug struct {
		EnableLogging bool `ini:"enable_logging"`
		DevDebug      bool `ini:"dev_debug"`
		LogToConsole  bool `ini:"log_to_console"`
	} `ini:"debug"`

	Sorting struct {
		EnableSorting bool `ini:"enable_sorting"`
		SortMode1     int  `ini:"sortmode1"`
		SortMode2     int  `ini:"sortmode2"`
		SortMode3     int  `ini:"sortmode3"`
		SortPlayer    bool `ini:"sort_player"`
	} `ini:"sorting"`

	Limits struct {
		MaximumAncillaries int `ini:"maximum_ancillaries"`
		GuildCooldown      int `ini:"guild_cooldown"`
	} `ini:"limits"`

	Saving struct {
		PostBattleSaving bool `ini:"post_battle_saving"`
	} `ini:"saving"`

	Info struct {
		HideArmyInfo       bool `ini:"hide_army_info"`
		AIRaidNotification bool `ini:"ai_raid_notification"`
		WatchtowerRadius   int  `ini:"watchtower_radius"`
	} `ini:"info"`

	Scripts struct {
		NaturalDisasters   bool `ini:"natural_disasters"`
		RandomAaAiStart    bool `ini:"random_aa_ai_start"`
		MergeDolAmroth     bool `ini:"merge_dol_amroth"`
		RandomizedStart    bool `ini:"randomized_start"`
		ShatteredAlliances bool `ini:"shattered_alliances"`
		LastStandArmies    bool `ini:"last_stand_armies"`
	} `ini:"scripts"`

	Battle struct {
		NoDefaultSkirmish  bool `ini:"no_default_skirmish"`
		DefaultBattleSpeed int  `ini:"default_battle_speed"`
	} `ini:"battle"`

	Difficulty struct {
		AggressiveRebels bool `ini:"aggressive_rebels"`
		AIFreeGenerals   bool `ini:"ai_free_generals"`
	} `ini:"difficulty"`
}

func (configurator *Configurator) GetConfigFilePath(file string) string {
	fmt.Printf("[Config] Attempting to get config file path for %s\n", file)
	// prod
	exePath, err := os.Executable()
	if err == nil {
		exeDir := filepath.Dir(exePath)
		configPath := filepath.Join(exeDir, "config", file)
		if _, err := os.Stat(configPath); err == nil {
			fmt.Printf("[Config] Found config file in executable directory: %s\n", configPath)
			return configPath
		}
	}

	// dev
	cwd, err := os.Getwd()
	if err == nil {
		configPath := filepath.Join(cwd, "config", file)
		if _, err := os.Stat(configPath); err == nil {
			fmt.Printf("[Config] Found config file in working directory: %s\n", configPath)
			return configPath
		}
	}

	fmt.Printf("[Config] Config file not found in either executable or working directory\n")
	return ""
}

func (configurator *Configurator) LoadConfigFile(file string) *ini.File {
	fmt.Printf("[Config] Loading config file: %s\n", file)
	configFile := configurator.GetConfigFilePath(file)
	fmt.Printf("[Config] Opening config file: %v\n", configFile)
	cfg, err := ini.Load(configFile)
	if err != nil {
		fmt.Printf("[Config] Fail to read file: %v\n", err)
		os.Exit(1)
	}
	// configurator.PrintConfig(cfg)

	return cfg
}

func (configurator *Configurator) ParseConfig(cfgFile *ini.File, cfgStruct interface{}) {
	fmt.Printf("[Config] Parsing config file to struct\n")
	err := cfgFile.MapTo(cfgStruct)
	if err != nil {
		fmt.Printf("[Config] Failed to map config file to struct: %v\n", err)
		os.Exit(1)
	}
}

func (configurator *Configurator) PrintConfig(cfg *ini.File) {
	fmt.Printf("[Config] Printing config file contents\n")
	for _, section := range cfg.Sections() {
		fmt.Printf("\n[%v]\n", section.Name())
		for _, option := range section.Keys() {
			fmt.Printf("%v = %v\n", option.Name(), option.Value())
		}
	}
}

func (configurator *Configurator) WriteConfigToFile(file string, cfgFile *ini.File, cfgData interface{}) {
	fmt.Printf("[Config] Writing config to file: %s\n", file)
	// Get the right file path
	configPath := configurator.GetConfigFilePath(file)
	openedFile, err := os.Create(configPath)
	if err != nil {
		fmt.Printf("[Config] Failed to open config file for writing: %v\n", err)
		os.Exit(1)
	}

	// Update the ini pointer with our struct from memory
	err = cfgFile.ReflectFrom(cfgData)
	if err != nil {
		fmt.Printf("[Config] Failed to update ini file from struct: %v\n", err)
		os.Exit(1)
	}

	// Write the config back to the file
	_, err = cfgFile.WriteTo(openedFile)
	if err != nil {
		fmt.Printf("[Config] Failed to write config to file: %v\n", err)
		os.Exit(1)
	}

	defer openedFile.Close()
	fmt.Printf("[Config] Successfully wrote config to file: %s\n", configPath)
}