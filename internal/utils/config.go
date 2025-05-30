package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

const CONFIG_FILE = "resources/config/AGO.cfg"

func LoadConfig() *ini.File {
	log.Printf("Opening AGO config file: %v", CONFIG_FILE)
	cfg, err := ini.Load("resources/config/AGO.cfg")
	if err != nil {
		log.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return cfg
}
