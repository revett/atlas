package base

import (
	"fmt"

	"github.com/revett/sepias/internal/file"
	"github.com/rs/zerolog/log"
)

const templateDirectoryPath = "./templates"

// ValidateTemplatesExist checks that the templates directory exists, and that
// it contains all of the required template Markdown files.
func ValidateTemplatesExist() error {
	log.Info().Msg("validating template directory exists")

	if err := file.DirectoryOrFileExists(templateDirectoryPath); err != nil {
		return err
	}

	log.Info().Msg("validating required markdown templates exist")

	for _, t := range requiredTemplates() {
		p := fmt.Sprintf("%s/%s.md", templateDirectoryPath, t)

		if err := file.DirectoryOrFileExists(p); err != nil {
			return err
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
		"review",
		"scratch",
		"system",
	}
}
