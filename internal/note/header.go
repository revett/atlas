package note

import (
	"fmt"
	"time"

	"github.com/revett/sepias/internal/language"
)

const phraseChunkLength = 4

// TODO: add option to disable cspell.
func generateFrontmatterHeader() (string, error) {
	format := `---
// cspell:disable-next-line
id: %s
created: %d (%s)
---
`

	d, err := language.Dictionary()
	if err != nil {
		return "", fmt.Errorf("unable to load dictionary from file: %w", err)
	}

	id, err := language.RandomPhrase(d, phraseChunkLength)
	if err != nil {
		return "", fmt.Errorf("failed to generate random phrase: %w", err)
	}

	return fmt.Sprintf(
		format, id, time.Now().Unix(), time.Now().Format(time.RFC1123),
	), nil
}
