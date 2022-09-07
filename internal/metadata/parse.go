package metadata

import (
	"fmt"
	"os"

	"github.com/adrg/frontmatter"
)

// Parse decodes the values within the Front Matter block at the top of the note
// at the given filepath, and returns them as a metadata.Fields type.
func Parse(p string) (Fields, error) {
	var fields Fields

	r, err := os.Open(p) //nolint:gosec
	if err != nil {
		return Fields{}, fmt.Errorf("failed to open file '%s': %w", p, err)
	}

	_, err = frontmatter.Parse(r, &fields)
	if err != nil {
		return Fields{}, fmt.Errorf("failed to parse front matter fields: %w", err)
	}

	return fields, nil
}
