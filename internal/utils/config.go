package config

import (
	agoConfig "ago-launcher/model"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

func getConfigFilePath() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)
	configPath := filepath.Join(exeDir, "resources", "config", "AGO.cfg")
	return configPath
}

func LoadConfigFile() *ini.File {
	configFile := getConfigFilePath()
	log.Printf("Opening AGO config file: %v", configFile)
	cfg, err := ini.Load(configFile)
	if err != nil {
		log.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	PrintConfig(cfg)
	return cfg
}

func ParseConfig(cfgFile *ini.File) agoConfig.AGOConfig {
	log.Printf("Parsing config file to struct")
	var cfgStruct agoConfig.AGOConfig
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
