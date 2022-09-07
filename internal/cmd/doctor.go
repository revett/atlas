package cmd

import (
	"fmt"
	"os"

	"github.com/revett/sepia/internal/base"
	"github.com/revett/sepia/internal/input"
	"github.com/revett/sepia/internal/metadata"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Doctor returns a cobra.Command type that checks that each note within the
// knowledge base is valid, using a number of different validation checks.
func Doctor() *cobra.Command {
	return &cobra.Command{
		Use:                   "doctor",
		Short:                 "Validates existing notes",
		DisableFlagsInUseLine: true,
		Args:                  cobra.NoArgs,
		RunE:                  doctorRunE,
	}
}

func doctorRunE(c *cobra.Command, args []string) error {
	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get current working directory: %w", err)
	}

	log.Info().Str("path", path).Msg("validating knowledge base")

	files, err := base.Read(path)
	if err != nil {
		return fmt.Errorf("failed to read notes in knowledge base: %w", err)
	}

	log.Info().Int("count", len(files)).Msgf("found %d notes", len(files))

	var foundErrors []error

	for _, file := range files {
		title := file.Name()

		if err := input.ValidateTitleFormat(title); err != nil {
			foundError := fmt.Errorf("./%s: invalid note title: %w", title, err)
			foundErrors = append(foundErrors, foundError)
		}

		if err := input.ValidateTitleBaseSchemaType(title); err != nil {
			foundError := fmt.Errorf("./%s: invalid base schema type: %w", title, err)
			foundErrors = append(foundErrors, foundError)
		}

		metaFields, err := metadata.Parse(title)
		if err != nil {
			foundError := fmt.Errorf(
				"./%s: unable to parse front matter metadata: %w", title, err,
			)
			foundErrors = append(foundErrors, foundError)
		}

		if err := metaFields.Validate(); err != nil {
			foundError := fmt.Errorf(
				"./%s: invalid front matter metadata: %w", title, err,
			)
			foundErrors = append(foundErrors, foundError)
		}
	}

	if err := base.ValidateTemplatesExist(); err != nil {
		foundError := fmt.Errorf("missing template files: %w", err)
		foundErrors = append(foundErrors, foundError)
	}

	log.Info().
		Int("count", len(foundErrors)).
		Msgf("detected %d validation errors", len(foundErrors))

	for _, foundError := range foundErrors {
		log.Warn().Msg(foundError.Error())
	}

	return nil
}
