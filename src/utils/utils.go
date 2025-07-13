package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
)

func RandomElement[T any](slice []T) (T, error) {
	var zero T
	if len(slice) == 0 {
		return zero, errors.New("no elements available")
	}
	return slice[rand.Intn(len(slice))], nil
}


// Run executable in the same directory as the Go program
func RunExecutable(exeName string) {
	fmt.Println("=== Running Local Executable ===")
	
	// Get the directory of the current executable
	execDir, err := os.Executable()
	if err != nil {
		fmt.Printf("Error getting executable path: %v\n", err)
		return
	}
	
	// Get the directory containing the executable
	dir := filepath.Dir(execDir)
	
	// Path to your executable (adjust the name as needed)
	exePath := filepath.Join(dir, exeName)
	
	fmt.Printf("Attempting to run: %s\n", exePath)
	
	// Check if file exists
	if _, err := os.Stat(exePath); os.IsNotExist(err) {
		fmt.Printf("Executable not found: %s\n", exePath)
		return
	}
	
	// Run the executable
	cmd := exec.Command(exePath)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running executable: %v\n", err)
	} else {
		fmt.Println("Executable completed successfully")
	}
}
