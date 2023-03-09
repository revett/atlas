package metadata

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/revett/atlas/internal/language"
)

//go:embed template.md
var template string

// Generate returns a string repesentation of the Front Matter metadata block to be included within
// a new note file.
func Generate() string {
	return fmt.Sprintf(
		template,
		language.HumanReadableID(),
		time.Now().Format(time.RFC1123),
	)
}
