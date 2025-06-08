package updater

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
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
	Version             string `json:"version"`
	Latest              bool   `json:"latest"`
	Url                 string `json:"url"`
	SaveGameCompatitble bool   `json:"sgc"`
}

func (updater *Updater) GetCurrentModVersion() {
	fmt.Println("[Updater] Retrieving mod version")
	jsonFile, err := os.Open("resources/uiCfg.json")
	if err != nil {
		fmt.Println("[Updater] could not open uiCfg file:", err)
		return
	}
	defer jsonFile.Close()

	jsonContent, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("[Updater] could not read uiCfg file:", err)
		return
	}

	modVersion := gjson.Get(string(jsonContent), "modVersion")
	fmt.Println("[Updater] Mod version", modVersion, "found")

	updater.CurrentVersion = ModVersion{
		Version: modVersion.String(),
	}
}

func (updater *Updater) GetLatestModVersion() {
	fmt.Println("[Updater] Retrieving latest mod version")
	// Local
	jsonFile, err := os.Open("resources/modVersions.json")
	if err != nil {
		fmt.Println("[Updater] could not open modVersions file:", err)
		return
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
		fmt.Println("[Updater] could not read modVersions file:", err)
	}
	var modVersions ModVersions

	jsonErr := json.Unmarshal(byteValue, &modVersions)
	if jsonErr != nil {
		fmt.Println("[Updater] could not unmarshal file:", jsonErr)
	}

	for i := 0; i < len(modVersions.ModVersions); i++ {
		version := modVersions.ModVersions[i]
		if version.Latest {
			fmt.Println("[Updater] Found latest version", version.Version)
			updater.LatestVersion = version
		}
	}
	updater.AvailableVersions = modVersions
}

func (updater *Updater) CheckForUpdate() (ModVersion, bool, error) {
	fmt.Println("[Updater] Checking for updates...")
	if updater.CurrentVersion.Version == "" {
		updater.GetCurrentModVersion()
	}

	updater.GetLatestModVersion()
	latestVersion := updater.LatestVersion

	if updater.LatestVersion.Version != updater.CurrentVersion.Version {
		fmt.Println("[Updater] !!! New mod version found !!!")
		fmt.Println("[Updater] Current Version: ", updater.CurrentVersion.Version)
		fmt.Println("[Updater] Latest Version: ", latestVersion.Version)
		updater.UpdateAvailable = true
		updater.LatestVersion = latestVersion
		return latestVersion, true, nil
	}

	fmt.Println("[Updater] No update available.")
	return updater.CurrentVersion, false, nil
}

func (updater *Updater) DownloadFile(url, dest string) error {
	fmt.Printf("[Updater] Downloading file from %s to %s\n", url, dest)
	req, _ := grab.NewRequest(dest, url)
	client := grab.NewClient()
	resp := client.Do(req)

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			fmt.Printf("[Updater]   transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())
		case <-resp.Done:
			if err := resp.Err(); err != nil {
				fmt.Println("[Updater] Download failed:", err)
				return err
			}
			fmt.Println("[Updater] Download saved to", resp.Filename)
			return nil
		}
	}
}

func (updater *Updater) GetUpdatesToApply() ([]ModVersion, error) {
	fmt.Println("[Updater] Determining updates to apply...")
	current, err := semver.NewVersion(updater.CurrentVersion.Version)
	if err != nil {
		fmt.Println("[Updater] Invalid current version:", updater.CurrentVersion.Version)
		return nil, fmt.Errorf("invalid current version: %w", err)
	}
	var updates []ModVersion
	for _, v := range updater.AvailableVersions.ModVersions {
		ver, err := semver.NewVersion(v.Version)
		if err != nil {
			fmt.Printf("[Updater] Skipping invalid version: %s\n", v.Version)
			continue // skip invalid versions
		}
		if ver.GreaterThan(current) {
			fmt.Printf("[Updater] Update available: %s\n", v.Version)
			updates = append(updates, v)
		}
	}
	// Sort updates by version ascending
	sort.Slice(updates, func(i, j int) bool {
		vi, _ := semver.NewVersion(updates[i].Version)
		vj, _ := semver.NewVersion(updates[j].Version)
		return vi.LessThan(vj)
	})
	fmt.Printf("[Updater] %d updates to apply.\n", len(updates))
	return updates, nil
}

// Applies all updates in order, updating the current version after each
func (updater *Updater) ApplyUpdatesSequentially(destDir string, onProgress func(idx, total int, v ModVersion)) error {
	fmt.Println("[Updater] Applying updates sequentially...")
	updates, err := updater.GetUpdatesToApply()
	if err != nil {
		fmt.Println("[Updater] Error getting updates to apply:", err)
		return err
	}
	total := len(updates)
	for i, update := range updates {
		if onProgress != nil {
			onProgress(i+1, total, update)
		}
		fmt.Printf("[Updater] Applying update %s (%d/%d)...\n", update.Version, i+1, total)
		err := updater.DownloadAndExtractUpdate(update, destDir)
		if err != nil {
			fmt.Printf("[Updater] Failed to apply update %s: %v\n", update.Version, err)
			return fmt.Errorf("failed to apply update %s: %w", update.Version, err)
		}
		updater.CurrentVersion = update // update current version after each
		fmt.Printf("[Updater] Update %s applied successfully.\n", update.Version)
	}
	fmt.Println("[Updater] All updates applied.")
	return nil
}

// DownloadAndExtractUpdate downloads and extracts a zip update, replacing files
func (updater *Updater) DownloadAndExtractUpdate(version ModVersion, destDir string) error {
	fmt.Printf("[Updater] Downloading and extracting update %s...\n", version.Version)
	tmpFile, err := os.CreateTemp("", "update-*.zip")
	if err != nil {
		fmt.Println("[Updater] Could not create temp file:", err)
		return err
	}
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())
	// Download
	err = updater.DownloadFile(version.Url, tmpFile.Name())
	if err != nil {
		fmt.Println("[Updater] Download failed:", err)
		return err
	}
	// Extract
	fmt.Printf("[Updater] Extracting %s to %s\n", tmpFile.Name(), destDir)
	return ExtractZip(tmpFile.Name(), destDir)
}

// ExtractZip extracts a zip archive to the destination directory, replacing files
func ExtractZip(src, dest string) error {
	fmt.Printf("[Updater] Extracting zip file %s to %s\n", src, dest)
	r, err := zip.OpenReader(src)
	if err != nil {
		fmt.Println("[Updater] Could not open zip file:", err)
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			fmt.Printf("[Updater] Creating directory: %s\n", fpath)
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			fmt.Printf("[Updater] Could not create directory for file %s: %v\n", fpath, err)
			return err
		}
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			fmt.Printf("[Updater] Could not open file for writing: %s: %v\n", fpath, err)
			return err
		}
		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			fmt.Printf("[Updater] Could not open file in zip: %s: %v\n", f.Name, err)
			return err
		}
		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			fmt.Printf("[Updater] Error copying file: %s: %v\n", fpath, err)
			return err
		}
		fmt.Printf("[Updater] Extracted file: %s\n", fpath)
	}
	fmt.Println("[Updater] Extraction complete.")
	return nil
}
