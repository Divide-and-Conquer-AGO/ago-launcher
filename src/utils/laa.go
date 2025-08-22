package utils

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

// IsLargeAddressAware checks if the given PE file has the Large Address Aware flag set.
func IsLargeAddressAware(fileName string) (bool, error) {
	exePath, err := os.Executable()
	if err != nil {
		return false, fmt.Errorf("could not get executable path: %w", err)
	}
	exeDir := filepath.Dir(exePath)
	filePath := filepath.Join(exeDir, fileName)

	Logger().Printf("[LAA] Checking for LAA: %s", filePath)
	var result bool
	err = prepareStream(filePath, func(f *os.File) error {
		var value uint16
		if err := binary.Read(f, binary.LittleEndian, &value); err != nil {
			return err
		}
		result = (value & 0x20) != 0
		return nil
	})
	Logger().Printf("[LAA] Checked: %s", filePath)
	return result, err
}

// SetLargeAddressAware sets the Large Address Aware flag if not already set.
func SetLargeAddressAware(fileName string) error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("could not get executable path: %w", err)
	}
	exeDir := filepath.Dir(exePath)
	filePath := filepath.Join(exeDir, fileName)

	Logger().Printf("[LAA] Applying LAA: %s", filePath)
	return prepareStream(filePath, func(f *os.File) error {
		var value uint16
		if err := binary.Read(f, binary.LittleEndian, &value); err != nil {
			return err
		}
		if (value & 0x20) == 0 {
			value |= 0x20
			// Move back 2 bytes to overwrite
			if _, err := f.Seek(-2, os.SEEK_CUR); err != nil {
				return err
			}
			if err := binary.Write(f, binary.LittleEndian, value); err != nil {
				return err
			}
			Logger().Printf("[LAA] Applied LAA successfully!")
		} else {
			Logger().Printf("[LAA] LAA already set.")
		}
		return nil
	})
}

// prepareStream handles opening the PE file and seeking to the Characteristics field in the COFF header.
func prepareStream(filePath string, action func(*os.File) error) error {
	f, err := os.OpenFile(filePath, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if fi.Size() < 0x3C {
		return fmt.Errorf("file too small to be a PE")
	}

	// Check MZ header
	var mz uint16
	if err := binary.Read(f, binary.LittleEndian, &mz); err != nil {
		return err
	}
	if mz != 0x5A4D {
		return fmt.Errorf("not a valid MZ executable")
	}

	// Go to PE header offset (at 0x3C)
	if _, err := f.Seek(0x3C, os.SEEK_SET); err != nil {
		return err
	}
	var peHeaderOffset int32
	if err := binary.Read(f, binary.LittleEndian, &peHeaderOffset); err != nil {
		return err
	}

	// Seek to PE header
	if _, err := f.Seek(int64(peHeaderOffset), os.SEEK_SET); err != nil {
		return err
	}

	// Check PE signature
	var peSig uint32
	if err := binary.Read(f, binary.LittleEndian, &peSig); err != nil {
		return err
	}
	if peSig != 0x4550 {
		return fmt.Errorf("not a valid PE file")
	}

	// Skip 0x12 bytes into the COFF header to reach Characteristics
	if _, err := f.Seek(0x12, os.SEEK_CUR); err != nil {
		return err
	}

	return action(f)
}
