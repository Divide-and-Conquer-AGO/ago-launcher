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
	LoggerVar  *log.Logger
	logFilePtr *os.File
)

// Called automatically when the application starts
func init() {
	var err error
	logFilePtr, err = os.OpenFile("ago-launcher.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
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

	// Get the directory of the current executable
	execDir, err := os.Executable()
	if err != nil {
		LoggerVar.Printf("Error getting executable path: %v\n", err)
		return
	}

	// Get the directory containing the executable
	dir := filepath.Dir(execDir)

	// Path to your executable (adjust the name as needed)
	exePath := filepath.Join(dir, exeName)

	LoggerVar.Printf("Attempting to run: %s\n", exePath)

	// Check if file exists
	if _, err := os.Stat(exePath); os.IsNotExist(err) {
		LoggerVar.Printf("Executable not found: %s\n", exePath)
		return
	}

	// Run the executable
	cmd := exec.Command(exePath)
	err = cmd.Run()
	if err != nil {
		LoggerVar.Printf("Error running executable: %v\n", err)
	} else {
		LoggerVar.Println("Executable completed successfully")
	}
}
