package updater

import (
	"ago-launcher/utils"
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/cavaliergopher/grab/v3"
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
	Version            string `json:"version"`
	Latest             bool   `json:"latest"`
	Url                string `json:"url"`
	SaveGameCompatible bool   `json:"sgc"`
}

func (updater *Updater) GetCurrentModVersion() {
	utils.Logger().Println("[Updater] Retrieving mod version")

	exePath, err := os.Executable()
	if err != nil {
		utils.Logger().Printf("[Updater] Could not get executable path: %v\n", err)
		return
	}
	exeDir := filepath.Dir(exePath)
	cfgPath := filepath.Join(exeDir, "eopData", "config", "uiCfg.json")

	jsonFile, err := os.Open(cfgPath)
	if err != nil {
		utils.Logger().Printf("[Updater] Could not open uiCfg file at %s: %v\n", cfgPath, err)
		return
	}
	defer jsonFile.Close()

	jsonContent, err := io.ReadAll(jsonFile)
	if err != nil {
		utils.Logger().Println("[Updater] could not read uiCfg file:", err)
		return
	}

	modVersion := gjson.Get(string(jsonContent), "modVersion")
	utils.Logger().Println("[Updater] Mod version", modVersion, "found")

	updater.CurrentVersion = ModVersion{
		Version: modVersion.String(),
	}
}

func (updater *Updater) GetLatestModVersion() {
	utils.Logger().Println("[Updater] Retrieving latest mod version")
	// // Local
	// jsonFile, err := os.Open("resources/modVersions.json")
	// if err != nil {
	// 	utils.Logger().Println("[Updater] could not open modVersions file:", err)
	// 	return
	// }
	// defer jsonFile.Close()

	// Remote
	resp, err := http.Get("https://raw.githubusercontent.com/Divide-and-Conquer-AGO/ago-launcher/refs/heads/main/src/resources/modVersions.json")
	if err != nil {
		utils.Logger().Println("could not fetch modVersions file from GitHub")
		return
	}
	defer resp.Body.Close()
	jsonFile := resp.Body

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		utils.Logger().Println("[Updater] could not read modVersions file:", err)
	}
	var modVersions ModVersions

	jsonErr := json.Unmarshal(byteValue, &modVersions)
	if jsonErr != nil {
		utils.Logger().Println("[Updater] could not unmarshal file:", jsonErr)
	}

	for i := 0; i < len(modVersions.ModVersions); i++ {
		version := modVersions.ModVersions[i]
		if version.Latest {
			utils.Logger().Println("[Updater] Found latest version", version.Version)
			updater.LatestVersion = version
		}
	}
	updater.AvailableVersions = modVersions
}

func (updater *Updater) CheckForUpdate() (ModVersion, bool, error) {
	utils.Logger().Println("[Updater] Checking for updates...")
	if updater.CurrentVersion.Version == "" {
		updater.GetCurrentModVersion()
	}

	updater.GetLatestModVersion()
	latestVersion := updater.LatestVersion

	if updater.LatestVersion.Version != updater.CurrentVersion.Version {
		utils.Logger().Println("[Updater] !!! New mod version found !!!")
		utils.Logger().Println("[Updater] Current Version: ", updater.CurrentVersion.Version)
		utils.Logger().Println("[Updater] Latest Version: ", latestVersion.Version)
		updater.UpdateAvailable = true
		updater.LatestVersion = latestVersion
		return latestVersion, true, nil
	}

	utils.Logger().Println("[Updater] No update available.")
	return updater.CurrentVersion, false, nil
}

func (updater *Updater) DownloadFile(url, dest string) error {
	utils.Logger().Printf("[Updater] Downloading file from %s to %s\n", url, dest)
	req, _ := grab.NewRequest(dest, url)
	client := grab.NewClient()
	resp := client.Do(req)

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			utils.Logger().Printf("[Updater]   transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())
		case <-resp.Done:
			if err := resp.Err(); err != nil {
				utils.Logger().Println("[Updater] Download failed:", err)
				return err
			}
			utils.Logger().Println("[Updater] Download saved to", resp.Filename)
			return nil
		}
	}
}

func (updater *Updater) GetUpdatesToApply() ([]ModVersion, error) {
	utils.Logger().Println("[Updater] Determining updates to apply...")
	current, err := semver.NewVersion(updater.CurrentVersion.Version)
	if err != nil {
		utils.Logger().Println("[Updater] Invalid current version:", updater.CurrentVersion.Version)
		return nil, fmt.Errorf("invalid current version: %w", err)
	}
	var updates []ModVersion
	for _, v := range updater.AvailableVersions.ModVersions {
		ver, err := semver.NewVersion(v.Version)
		if err != nil {
			utils.Logger().Printf("[Updater] Skipping invalid version: %s\n", v.Version)
			continue // skip invalid versions
		}
		if ver.GreaterThan(current) {
			utils.Logger().Printf("[Updater] Update available: %s\n", v.Version)
			updates = append(updates, v)
		}
	}
	// Sort updates by version ascending
	sort.Slice(updates, func(i, j int) bool {
		vi, _ := semver.NewVersion(updates[i].Version)
		vj, _ := semver.NewVersion(updates[j].Version)
		return vi.LessThan(vj)
	})
	utils.Logger().Printf("[Updater] %d updates to apply.\n", len(updates))
	return updates, nil
}

// Applies all updates in order, updating the current version after each
func (updater *Updater) ApplyUpdatesSequentially(destDir string, onProgress func(idx, total int, v ModVersion)) error {
	utils.Logger().Println("[Updater] Applying updates sequentially...")
	updates, err := updater.GetUpdatesToApply()
	if err != nil {
		utils.Logger().Println("[Updater] Error getting updates to apply:", err)
		return err
	}
	total := len(updates)
	for i, update := range updates {
		if onProgress != nil {
			onProgress(i+1, total, update)
		}
		utils.Logger().Printf("[Updater] Applying update %s (%d/%d)...\n", update.Version, i+1, total)
		err := updater.DownloadAndExtractUpdate(update, destDir)
		if err != nil {
			utils.Logger().Printf("[Updater] Failed to apply update %s: %v\n", update.Version, err)
			return fmt.Errorf("failed to apply update %s: %w", update.Version, err)
		}
		updater.CurrentVersion = update // update current version after each
		utils.Logger().Printf("[Updater] Update %s applied successfully.\n", update.Version)
	}
	utils.Logger().Println("[Updater] All updates applied.")
	return nil
}

// DownloadAndExtractUpdate downloads and extracts a zip update, replacing files
func (updater *Updater) DownloadAndExtractUpdate(version ModVersion, destDir string) error {
	utils.Logger().Printf("[Updater] Downloading and extracting update %s...\n", version.Version)

	utils.Logger().Printf("[Updater] Target directory: %s\n", destDir)

	tmpFile, err := os.CreateTemp("", "update-*.zip")
	if err != nil {
		utils.Logger().Println("[Updater] Could not create temp file:", err)
		return err
	}
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	// Download
	err = updater.DownloadFile(version.Url, tmpFile.Name())
	if err != nil {
		utils.Logger().Println("[Updater] Download failed:", err)
		return err
	}

	// Extract
	utils.Logger().Printf("[Updater] Extracting %s to %s\n", tmpFile.Name(), destDir)
	return ExtractZip(tmpFile.Name(), destDir)
}

// ExtractZip extracts a zip archive to the destination directory, replacing files
func ExtractZip(src, dest string) error {
	utils.Logger().Printf("[Updater] Extracting zip file %s to %s\n", src, dest)
	r, err := zip.OpenReader(src)
	if err != nil {
		utils.Logger().Println("[Updater] Could not open zip file:", err)
		return err
	}
	defer r.Close()

	// Ensure destination directory exists
	if err := os.MkdirAll(dest, os.ModePerm); err != nil {
		utils.Logger().Printf("[Updater] Could not create destination directory %s: %v\n", dest, err)
		return err
	}

	// Clean the destination path once
	cleanDest := filepath.Clean(dest)

	for _, f := range r.File {
		// Clean the file name to handle different path separators
		cleanName := filepath.Clean(f.Name)
		fpath := filepath.Join(cleanDest, cleanName)

		// Security check: prevent zip slip attacks
		// Use Rel to check if the path is within the destination
		rel, err := filepath.Rel(cleanDest, fpath)
		if err != nil || len(rel) > 0 && rel[0] == '.' && rel[1] == '.' {
			utils.Logger().Printf("[Updater] Skipping invalid file path in zip: %s (would extract to: %s)\n", f.Name, fpath)
			continue
		}

		utils.Logger().Printf("[Updater] Processing: %s -> %s\n", f.Name, fpath)

		if f.FileInfo().IsDir() {
			utils.Logger().Printf("[Updater] Creating directory: %s\n", fpath)
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			utils.Logger().Printf("[Updater] Could not create directory for file %s: %v\n", fpath, err)
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			utils.Logger().Printf("[Updater] Could not open file for writing: %s: %v\n", fpath, err)
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			utils.Logger().Printf("[Updater] Could not open file in zip: %s: %v\n", f.Name, err)
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			utils.Logger().Printf("[Updater] Error copying file: %s: %v\n", fpath, err)
			return err
		}
		utils.Logger().Printf("[Updater] Extracted file: %s\n", fpath)
	}
	utils.Logger().Println("[Updater] Extraction complete.")
	return nil
}
