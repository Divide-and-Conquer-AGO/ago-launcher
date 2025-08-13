package utils

import (
	"errors"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"

	"fyne.io/fyne/v2"
)

var (
	LoggerVar *log.Logger
)

func init() {
	setupLogger()
}

func Logger() *log.Logger {
	return LoggerVar
}

func setupLogger() {
	exePath, err := os.Executable()
	var logFilePtr *os.File
	if err == nil {
		logFilePath := filepath.Join(filepath.Dir(exePath), "AGO_Launcher.log")
		logFilePtr, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	}
	if err != nil {
		log.Printf("Failed to open log file: %v\n", err)
		LoggerVar = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		LoggerVar = log.New(logFilePtr, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
	if exePath != "" {
		dir := filepath.Dir(exePath)
		Logger().Println("[Utils] Running from directory:", dir)
	}
}

func RandomElement[T any](slice []T) (T, error) {
	var zero T
	if len(slice) == 0 {
		return zero, errors.New("no elements available")
	}
	return slice[rand.Intn(len(slice))], nil
}

// Run executable relative to the location of the running executable
func RunExecutable(exeName string) {
    var app = fyne.CurrentApp()
    Logger().Println("=== Running Local Executable ===")

    exePath, err := os.Executable()
    if err != nil {
        Logger().Printf("Error getting executable path: %v\n", err)
        return
    }

    exeDir := filepath.Dir(exePath)
    targetExePath := filepath.Join(exeDir, exeName)
    
    // Clean the path to remove any malformed separators
    targetExePath = filepath.Clean(targetExePath)

    Logger().Printf("[Utils] Attempting to run: %s\n", targetExePath)
    Logger().Printf("[Utils] Executable directory: %s\n", exeDir)
    Logger().Printf("[Utils] Target executable name: %s\n", exeName)

    // Check if file exists
    if _, err := os.Stat(targetExePath); os.IsNotExist(err) {
        Logger().Printf("[Utils] Executable not found: %s\n", targetExePath)
        return
    }

    // Use proper escaping for paths with spaces
    cmd := exec.Command("cmd", "/c", "start", "/D", exeDir, "", exeName)
    err = cmd.Start()
    if err != nil {
        var msg = "Error running executable: " + err.Error()
        Logger().Println(msg)
        app.SendNotification(&fyne.Notification{
            Title:   "Permission Error", 
            Content: "Try running launcher as Administrator or check antivirus settings",
        })
        Logger().Println("[Utils] *** PERMISSION ERROR - Try running as Administrator or check antivirus settings ***")
    } else {
        Logger().Println("[Utils] Executable started successfully")
        app.Quit()
    }
}