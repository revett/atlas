package metadata

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/revett/sepia/internal/language"
)

//go:embed template.md
var template string

// Generate returns a string repesentation of the Front Matter metadata block
// to be included within a new note file.
func Generate() (string, error) {
	d, err := language.Dictionary()
	if err != nil {
		return "", fmt.Errorf("unable to load dictionary from file: %w", err)
	}

	id, err := language.RandomPhrase(d, idFieldChunks)
	if err != nil {
		return "", fmt.Errorf("failed to generate random phrase: %w", err)
	}

	return fmt.Sprintf(template, id, time.Now().Format(time.RFC1123)), nil
}
