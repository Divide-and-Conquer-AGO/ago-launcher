package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type Configurator struct {
	AGOConfigFile *ini.File
	ModConfigFile *ini.File

	AGOConfig    	AGOConfig
	ModConfig    	ModConfig
}

type ModConfig struct {
	Video struct {
		BorderlessWindow bool `ini:"borderless_window"`
		Windowed      bool `ini:"windowed"`
		BattleResolution  string `ini:"battle_resolution"`
		CampaignResolution  string `ini:"campaign_resolution"`
		Bloom      bool `ini:"bloom"`
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
    // prod
    exePath, err := os.Executable()
    if err == nil {
        exeDir := filepath.Dir(exePath)
        configPath := filepath.Join(exeDir, "config", file)
        if _, err := os.Stat(configPath); err == nil {
            return configPath
        }
    }

    // dev
    cwd, err := os.Getwd()
    if err == nil {
        configPath := filepath.Join(cwd, "config",file)
        if _, err := os.Stat(configPath); err == nil {
            return configPath
        }
    }

    log.Fatalf("cfg file not found in either executable or working directory")
    return ""
}

func (configurator *Configurator) LoadConfigFile(file string) *ini.File {
	configFile := configurator.GetConfigFilePath(file)
	log.Printf("Opening config file: %v", configFile)
	cfg, err := ini.Load(configFile)
	if err != nil {
		log.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	// configurator.PrintConfig(cfg)

	return cfg
}

func (configurator *Configurator) ParseConfig(cfgFile *ini.File, cfgStruct interface{}) {
	log.Printf("Parsing config file to struct")
	err := cfgFile.MapTo(cfgStruct)
	if err != nil {
		log.Fatalf("Failed to map config file to struct: %v", err)
	}
}

func (configurator *Configurator) PrintConfig(cfg *ini.File) {
	log.SetFlags(0)
	for _, section := range cfg.Sections() {
		log.Printf("\n[%v]", section.Name())
		for _, option := range section.Keys() {
			log.Printf("%v = %v", option.Name(), option.Value())
		}
	}
}

func (configurator *Configurator) WriteConfigToFile(file string, cfgFile *ini.File, cfgData interface{}) {
	// Get the right file path
	configPath := configurator.GetConfigFilePath(file)
	openedFile, err := os.Create(configPath)
	if err != nil {
		log.Fatalf("Failed to open config file for writing: %v", err)
	}

	// Update the ini pointer with our struct from memory
	err = cfgFile.ReflectFrom(cfgData)
	if err != nil {
		log.Fatalf("Failed to update ini file from struct: %v", err)
	}

	// Write the config back to the file
	_, err = cfgFile.WriteTo(openedFile)
	if err != nil {
		log.Fatalf("Failed to write config to file: %v", err)
	}

	defer openedFile.Close()
}