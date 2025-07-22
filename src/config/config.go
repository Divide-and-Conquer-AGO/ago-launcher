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
	AGOConfigFile *ini.File
	ModConfigFile *ini.File

	AGOConfig       AGOConfig
	ModConfig       ModConfig
	EOPConfig       EOPConfig
}

// Loads a file using json/ini format and loads it into a given struct
func (configurator *Configurator) LoadConfigFile(path string, cfgStruct interface{}) *ini.File {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("[Config] Could not get executable path: %v", err)
	}
	baseDir := filepath.Dir(exePath)
	absPath := path
	if !filepath.IsAbs(path) {
		absPath = filepath.Join(baseDir, path)
	}
	utils.Logger().Printf("[Config] Attempting to open config file at full path: %s\n", absPath)

	switch filepath.Ext(absPath) {
	case ".json":
		jsonFile, err := os.Open(absPath)
		if err != nil {
			log.Fatalf("[Config] Could not open json config file at %s: %v", absPath, err)
		}
		defer jsonFile.Close()

		byteValue, err := io.ReadAll(jsonFile)
		if err != nil {
			log.Fatalf("[Config] Fail to json read file at %s: %v\n", absPath, err)
		}
		utils.Logger().Printf("[Config] Parsing json config file at %s to struct\n", absPath)
		err = json.Unmarshal(byteValue, cfgStruct)
		if err != nil {
			log.Fatalf("[Config] Failed to map json config file at %s to struct: %v\n", absPath, err)
		}
		return nil
	case ".cfg":
		cfg, err := ini.LoadSources(ini.LoadOptions{
			AllowShadows:   true,
			SpaceBeforeInlineComment: false,
		}, absPath)
		if err != nil {
			log.Fatalf("[Config] Fail to read ini file at %s: %v\n", absPath, err)
		}
		utils.Logger().Printf("[Config] Parsing ini config file at %s to struct\n", absPath)
		err = cfg.MapTo(cfgStruct)
		if err != nil {
			log.Fatalf("[Config] Failed to map ini config file at %s to struct: %v\n", absPath, err)
		}
		return cfg
	}
	return nil
}

// Writes the config data back to the json/ini file
func (configurator *Configurator) WriteConfigToFile(path string, cfgData interface{}, filePtr *ini.File) {
	utils.Logger().Printf("[Config] Writing config to file: %s\n", path)

	// Get the executable directory
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("[Config] Could not get executable path: %v", err)
	}
	baseDir := filepath.Dir(exePath)
	absPath := path
	if !filepath.IsAbs(path) {
		absPath = filepath.Join(baseDir, path)
	}

	// Open the file for writing
	openedFile, err := os.Create(absPath)
	if err != nil {
		log.Fatalf("[Config] Failed to open config file for writing: %v\n", err)
		os.Exit(1)
	}
	defer openedFile.Close()

	switch filepath.Ext(absPath) {
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

	utils.Logger().Printf("[Config] Successfully wrote config to file: %s\n", absPath)
}

func (configurator *Configurator) LoadAllConfigFiles() {
	// AGO.cfg
	agoCfgPtr := configurator.LoadConfigFile("AGO.cfg", &configurator.AGOConfig)
	configurator.AGOConfigFile = agoCfgPtr

	// TATW.cfg
	modCfgPtr := configurator.LoadConfigFile("TATW.cfg", &configurator.ModConfig)
	configurator.ModConfigFile = modCfgPtr

	// gameCfg.json
	configurator.LoadConfigFile("eopData/config/gameCfg.json", &configurator.EOPConfig.GameCfg)

	// battleCfg.json
	configurator.LoadConfigFile("eopData/config/battlesCfg.json", &configurator.EOPConfig.BattlesCfg)
}
