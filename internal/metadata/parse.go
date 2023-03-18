package metadata

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/adrg/frontmatter"
	"github.com/revett/atlas/internal/config"
)

// Parse decodes the values within the Front Matter block at the top of the note
// at the given filepath, and returns them as a metadata.Fields type.
func Parse(cfg config.Config, filename string) (Fields, error) {
	var fields Fields

	path := filepath.ToSlash(
		path.Join(cfg.Path, filename),
	)

	r, err := os.Open(path) //nolint:gosec
	if err != nil {
		return Fields{}, fmt.Errorf("opening file '%s': %w", path, err)
	}

	_, err = frontmatter.Parse(r, &fields)
	if err != nil {
		return Fields{}, fmt.Errorf("parsing front matter fields: %w", err)
	}

	return fields, nil
}
