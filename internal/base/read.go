package base

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Read reads the current directory, and returns a slice of all fs.DirEntry
// elements which are a file and have the filenmae extension of ".md".
func Read(path string) ([]fs.DirEntry, error) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("reading directory at '%s': %w", path, err)
	}

	notes := []fs.DirEntry{}

	for _, dirEntry := range dirEntries {
		if !dirEntry.Type().IsRegular() {
			continue
		}

		if filepath.Ext(dirEntry.Name()) != ".md" {
			continue
		}

		notes = append(notes, dirEntry)
	}

	return notes, nil
}
