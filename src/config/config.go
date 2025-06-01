package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

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

func GetConfigFilePath() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)
	configPath := filepath.Join(exeDir, "config", "AGO.cfg")
	return configPath
}

func LoadConfigFile() *ini.File {
	configFile := GetConfigFilePath()
	log.Printf("Opening AGO config file: %v", configFile)
	cfg, err := ini.Load(configFile)
	if err != nil {
		log.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	PrintConfig(cfg)
	return cfg
}

func ParseConfig(cfgFile *ini.File) AGOConfig {
	log.Printf("Parsing config file to struct")
	var cfgStruct AGOConfig
	err := cfgFile.MapTo(&cfgStruct)
	if err != nil {
		log.Fatalf("Failed to map config file to struct: %v", err)
	}
	return cfgStruct
}

func PrintConfig(cfg *ini.File) {
	log.SetFlags(0)
	for _, section := range cfg.Sections() {
		log.Printf("\n[%v]", section.Name())
		for _, option := range section.Keys() {
			log.Printf("%v = %v", option.Name(), option.Value())
		}
	}
}
