package base

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/revett/atlas/internal/config"
	"github.com/revett/atlas/internal/file"
)

// ValidateTemplatesExist checks that the templates directory exists, and that
// it contains all of the required template Markdown files.
func ValidateTemplatesExist(cfg config.Config) error {
	templatesDirectoryPath := filepath.ToSlash(
		path.Join(cfg.Path, "templates"),
	)

	if err := file.DirectoryOrFileExists(templatesDirectoryPath); err != nil {
		return fmt.Errorf(
			"checking if template directory exists: %w", err,
		)
	}

	for _, t := range requiredTemplates() {
		f := fmt.Sprintf("%s.md", t)
		path := filepath.ToSlash(
			path.Join(templatesDirectoryPath, f),
		)

		if err := file.DirectoryOrFileExists(path); err != nil {
			return fmt.Errorf(
				"checking if required template '%s' exists: %w", path, err,
			)
		}
	}

	return nil
}

func requiredTemplates() []string {
	return []string{
		"area",
		"entity",
		"meeting",
		"meeting.interview.cultural",
		"meeting.interview.technical",
		"project",
		"scratch",
		"system",
	}
}
