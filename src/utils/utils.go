package utils

import (
	"errors"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
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
        Logger().Println("Running from directory:", dir)
    }
}

func RandomElement[T any](slice []T) (T, error) {
	var zero T
	if len(slice) == 0 {
		return zero, errors.New("no elements available")
	}
	return slice[rand.Intn(len(slice))], nil
}

// Run executable relative to current working directory
func RunExecutable(exeName string) {
    LoggerVar.Println("=== Running Local Executable ===")

    // Use current working directory to find the executable
    cwd, err := os.Getwd()
    if err != nil {
        LoggerVar.Printf("Error getting current directory: %v\n", err)
        return
    }
    
    exePath := filepath.Join(cwd, exeName)
    LoggerVar.Printf("Attempting to run: %s\n", exePath)

    // Check if file exists
    if _, err := os.Stat(exePath); os.IsNotExist(err) {
        LoggerVar.Printf("Executable not found: %s\n", exePath)
        return
    }

    // Run the executable
    cmd := exec.Command(exePath)
    err = cmd.Start()
    if err != nil {
        LoggerVar.Printf("Error running executable: %v\n", err)
        LoggerVar.Println("*** PERMISSION ERROR - Try running as Administrator or check antivirus settings ***")
    } else {
        LoggerVar.Println("Executable started successfully")
    }
}