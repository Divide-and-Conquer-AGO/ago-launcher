package updater

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/tidwall/gjson"
)

type Updater struct {
	CurrentVersion    ModVersion
	LatestVersion     ModVersion
	AvailableVersions ModVersions
	UpdateAvailable   bool
}

type ModVersions struct {
	ModVersions []ModVersion `json:"modVersions"`
}

type ModVersion struct {
	Version             string `json:"version"`
	Latest              bool   `json:"latest"`
	Url                 string `json:"url"`
	SaveGameCompatitble bool   `json:"sgc"`
}

func (updater *Updater) GetCurrentModVersion() {
	fmt.Println("Retrieving mod version")
	jsonFile, err := os.Open("resources/uiCfg.json")
	if err != nil {
		fmt.Println("could not open uiCfg file")
		return
	}
	defer jsonFile.Close()

	jsonContent, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("could not read uiCfg file")
		return
	}

	modVersion := gjson.Get(string(jsonContent), "modVersion")
	fmt.Println("Mod version", modVersion, "found")

	updater.CurrentVersion = ModVersion{
		Version: modVersion.String(),
	}
}

func (updater *Updater) GetLatestModVersion() {
	// Local
	jsonFile, err := os.Open("resources/modVersions.json")
	if err != nil {
	}
	defer jsonFile.Close()

	// Remote
	// resp, err := http.Get("https://raw.githubusercontent.com/EddieEldridge/ago-launcher/refs/heads/main/src/resources/modVersions.json?token=<>")
	// if err != nil {
	// 	fmt.Println("could not fetch modVersions file from GitHub")
	// 	return
	// }
	// defer resp.Body.Close()
	// jsonFile := resp.Body

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("could not open modVersions file")
	}
	var modVersions ModVersions

	jsonErr := json.Unmarshal(byteValue, &modVersions)
	if jsonErr != nil {
		fmt.Println("could not unmarshal file")
	}

	for i := 0; i < len(modVersions.ModVersions); i++ {
		version := modVersions.ModVersions[i]
		if version.Latest {
			fmt.Println("Found latest version", version.Version)
			updater.LatestVersion = version
		}
	}
}

func (updater *Updater) CheckForUpdate() (ModVersion, bool, error) {
	if updater.CurrentVersion.Version == "" {
		updater.GetCurrentModVersion()
	}

	updater.GetLatestModVersion()
	latestVersion := updater.LatestVersion

	if updater.LatestVersion.Version != updater.CurrentVersion.Version {
		fmt.Println("!!! New mod version found !!!")
		fmt.Println("Current Version: ", updater.CurrentVersion.Version)
		fmt.Println("Latest Version: ", latestVersion.Version)
		updater.UpdateAvailable = true
		updater.LatestVersion = latestVersion
		return latestVersion, true, nil
	}

	return updater.CurrentVersion, false, nil
}
