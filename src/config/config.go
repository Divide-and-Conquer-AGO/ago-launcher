package config

import (
	"ago-launcher/utils"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type Configurator struct {
	AGOConfigFile   *ini.File
	ModConfigFile 	*ini.File

	AGOConfig       AGOConfig
	ModConfig       ModConfig
	EOPConfig       EOPConfig
	ConfigLocations []string
}

func (configurator *Configurator) GetConfigFilePath(file string) string {
	utils.Logger().Printf("[Config] Attempting to get config file path for %s and paths %s\n", file, configurator.ConfigLocations)

	for _, path := range configurator.ConfigLocations {
		// prod
		exePath, err := os.Executable()
		if err == nil {
			exeDir := filepath.Dir(exePath)
			configPath := filepath.Join(exeDir, path, file)
			if _, err := os.Stat(configPath); err == nil {
				utils.Logger().Printf("[Config] Found config file in executable directory: %s\n", configPath)
				return configPath
			}
		}

		// dev
		cwd, err := os.Getwd()
		if err == nil {
			configPath := filepath.Join(cwd, path, file)
			if _, err := os.Stat(configPath); err == nil {
				utils.Logger().Printf("[Config] Found config file in working directory: %s\n", configPath)
				return configPath
			}
		}
	}

	utils.Logger().Printf("[Config] Config file not found in either executable or working directory\n")
	return ""
}

// Loads a file using json/ini format and loads it into a given struct
func (configurator *Configurator) LoadConfigFile(file string, cfgStruct interface{}) *ini.File {
	utils.Logger().Printf("[Config] Loading config file: %s\n", file)

	configPath := configurator.GetConfigFilePath(file)
	utils.Logger().Printf("[Config] Opening config file: %v\n", configPath)

	switch filepath.Ext(configPath) {
		case ".json":
			jsonFile, err := os.Open(configPath)
			if err != nil {
				log.Fatalf("[Config] Could not open json config file: %v", err)
			}
			defer jsonFile.Close()

			byteValue, err := io.ReadAll(jsonFile)
			if err != nil {
				log.Fatalf("[Config] Fail to json read file: %v\n", err)
			}
			utils.Logger().Printf("[Config] Parsing json config file to struct\n")
			err = json.Unmarshal(byteValue, cfgStruct)
			if err != nil {
				log.Fatalf("[Config] Failed to map json config file to struct: %v\n", err)
			}
			return nil
		case ".cfg":
			cfg, err := ini.Load(configPath)

			if err != nil {
				log.Fatalf("[Config] Fail to read ini file: %v\n", err)
			}
			utils.Logger().Printf("[Config] Parsing ini config file to struct\n")
			err = cfg.MapTo(cfgStruct)
			if err != nil {
				log.Fatalf("[Config] Failed to map ini config file to struct: %v\n", err)
			}
			return cfg
	}
	return nil
}

// Writes the config data back to the json/ini file
func (configurator *Configurator) WriteConfigToFile(file string, cfgData interface{}, filePtr *ini.File) {
	utils.Logger().Printf("[Config] Writing config to file: %s\n", file)
	
	// Get the right file path
	configPath := configurator.GetConfigFilePath(file)
	openedFile, err := os.Create(configPath)
	if err != nil {
		log.Fatalf("[Config] Failed to open config file for writing: %v\n", err)
		os.Exit(1)
	}
	defer openedFile.Close()

	switch filepath.Ext(configPath) {
		case ".json":
			encoder := json.NewEncoder(openedFile) 
			encoder.SetIndent("", "    ")
			err := encoder.Encode(cfgData)
			if err != nil {
				log.Fatalf("[Config] Fail to write file: %v\n", err)
			}
		case ".cfg":
			err = filePtr.ReflectFrom(cfgData)
			if err != nil {
				log.Fatalf("[Config] Failed to update ini file from struct: %v\n", err)
			}

			// Write the config back to the file
			_, err = filePtr.WriteTo(openedFile)
			if err != nil {
				log.Fatalf("[Config] Failed to write config to file: %v\n", err)
			}
	}

	utils.Logger().Printf("[Config] Successfully wrote config to file: %s\n", configPath)
}

func (configurator *Configurator) LoadAllConfigFiles() {
	// AGO.cfg
	agoCfgPtr := configurator.LoadConfigFile("AGO.cfg", &configurator.AGOConfig)
	configurator.AGOConfigFile = agoCfgPtr

	// TATW.cfg
	modCfgPtr := configurator.LoadConfigFile("TATW.cfg", &configurator.ModConfig)
	configurator.ModConfigFile = modCfgPtr

	// gameCfg.json
	configurator.LoadConfigFile("gameCfg.json", &configurator.EOPConfig.GameCfg)

	// uiCfg.json
	configurator.LoadConfigFile("battlesCfg.json", &configurator.EOPConfig.BattlesCfg)
}
