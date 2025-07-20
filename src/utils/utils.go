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
    // For fyne package builds, use Fyne's logging which works properly
    if isFynePackaged() {
        // Use Fyne's logging system
        LoggerVar = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
    } else {
        // For go build, use your file logging
        setupFileLogging()
    }
}

func isFynePackaged() bool {
    // Simple check: if we're running from a temp directory or system directory,
    // we're likely packaged
    execPath, err := os.Executable()
    if err != nil {
        return false
    }
    return filepath.Dir(execPath) != filepath.Dir(os.Args[0])
}

func setupFileLogging() {
    execPath, err := os.Executable()
    if err == nil {
        execDir := filepath.Dir(execPath)
        os.Chdir(execDir)
    }
    
    logFilePtr, err := os.OpenFile("AGO_Launcher.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
    if err != nil {
        log.Printf("Failed to open log file: %v\n", err)
        LoggerVar = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
    } else {
        LoggerVar = log.New(logFilePtr, "", log.Ldate|log.Ltime|log.Lshortfile)
    }
}

func Logger() *log.Logger {
	return LoggerVar
}

func RandomElement[T any](slice []T) (T, error) {
	var zero T
	if len(slice) == 0 {
		return zero, errors.New("no elements available")
	}
	return slice[rand.Intn(len(slice))], nil
}

// Run executable in the same directory as the Go program
func RunExecutable(exeName string) {
    LoggerVar.Println("=== Running Local Executable ===")

    var exePath string
    
    if isFynePackaged() {
        // For fyne package, use executable directory
        execDir, err := os.Executable()
        if err != nil {
            LoggerVar.Printf("Error getting executable path: %v\n", err)
            return
        }
        dir := filepath.Dir(execDir)
        exePath = filepath.Join(dir, exeName)
    } else {
        // For go build, use current working directory
        cwd, err := os.Getwd()
        if err != nil {
            LoggerVar.Printf("Error getting current directory: %v\n", err)
            return
        }
        exePath = filepath.Join(cwd, exeName)
    }

    LoggerVar.Printf("Attempting to run: %s\n", exePath)

    // Check if file exists
    if _, err := os.Stat(exePath); os.IsNotExist(err) {
        LoggerVar.Printf("Executable not found: %s\n", exePath)
        return
    }

    // Run the executable
    cmd := exec.Command(exePath)
    err := cmd.Start()
    if err != nil {
        LoggerVar.Printf("Error running executable: %v\n", err)
        
        // Check if it's a permission-related error
		LoggerVar.Println("*** PERMISSION ERROR - Try running as Administrator or check antivirus settings ***")
		
		// Show Fyne toast notification
		fyneApp := fyne.CurrentApp()
		fyneApp.SendNotification(&fyne.Notification{
			Title:   "Permission Error",
			Content: "Failed to launch " + exeName + ". Try running this application as Administrator or check antivirus settings.",
		})
    } else {
        LoggerVar.Println("Executable started successfully")
    }
}