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
// Run executable relative to the location of the running executable
func RunExecutable(exeName string) {
    LoggerVar.Println("=== Running Local Executable ===")

    exePath, err := os.Executable()
    if err != nil {
        LoggerVar.Printf("Error getting executable path: %v\n", err)
        return
    }

    exeDir := filepath.Dir(exePath)
    targetExePath := filepath.Join(exeDir, exeName)
    LoggerVar.Printf("Attempting to run: %s\n", targetExePath)

    // Check if file exists
    if _, err := os.Stat(targetExePath); os.IsNotExist(err) {
        LoggerVar.Printf("Executable not found: %s\n", targetExePath)
        return
    }

    // Run the executable with working directory set to exeDir
    cmd := exec.Command(targetExePath)
    cmd.Dir = exeDir
    err = cmd.Start()
    if err != nil {
        LoggerVar.Printf("Error running executable: %v\n", err)
        LoggerVar.Println("*** PERMISSION ERROR - Try running as Administrator or check antivirus settings ***")
    } else {
        LoggerVar.Println("Executable started successfully")
    }
}