package file

import (
	"fmt"
	"os"
)

// DirectoryOrFileExists checks if a file or directory exists locally.
func DirectoryOrFileExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		return fmt.Errorf("file '%s' does not exist: %w", path, err)
	}

	return fmt.Errorf("checking if '%s' file exists: %w", path, err)
}
