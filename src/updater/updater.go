package updater

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/tidwall/gjson"
)

type Updater struct {
	ModVersions ModVersions
}

type ModVersions struct {
	ModVersions []ModVersion `json:"modVersions"`
}

type ModVersion struct {
	Version string `json:"version"`
	Latest  bool   `json:"latest"`
	Url     string `json:"url"`
}

func (updater *Updater) GetModVersion() (ModVersion, error) {
	fmt.Println("Retrieving mod version")
	jsonFile, err := os.Open("resources/uiCfg.json")
	if err != nil {
		fmt.Println("could not open uiCfg file")
		return ModVersion{}, err
	}
	defer jsonFile.Close()

	jsonContent, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("could not read uiCfg file")
		return ModVersion{}, err
	}

	modVersion := gjson.Get(string(jsonContent), "modVersion")
	fmt.Println("Mod version", modVersion, "found")

	return ModVersion{}, nil
}

func (updater *Updater) GetLatestModVersion() (ModVersion, error) {
	jsonFile, err := os.Open("resources/uiCfg.json")
	if err != nil {
		return ModVersion{}, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return ModVersion{}, err
	}
	var modVersions ModVersions

	jsonErr := json.Unmarshal(byteValue, &modVersions)
	if jsonErr != nil {
		return ModVersion{}, jsonErr
	}

	var latestVersion ModVersion
	for i := 0; i < len(modVersions.ModVersions); i++ {
		version := modVersions.ModVersions[i]
		if version.Latest {
			fmt.Println("found latest version", version.Version)
			return latestVersion, nil
		}
	}
	return ModVersion{}, err
}

func (updater *Updater) CheckForUpdate() {

}
