package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/revett/sepias/internal/base"
	"github.com/revett/sepias/internal/input"
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

	notes, err := base.Read(path)
	if err != nil {
		return fmt.Errorf("failed to read notes in knowledge base: %w", err)
	}

	log.Info().Int("count", len(notes)).Msgf("found %d notes", len(notes))

	var foundErrors []error

	for _, note := range notes {
		withoutExtension := strings.TrimSuffix(
			note.Name(), filepath.Ext(note.Name()),
		)

		if err := input.ValidateTitleFormat(withoutExtension); err != nil {
			foundError := fmt.Errorf(
				"invalid note title found in '%s' file: %w", note.Name(), err,
			)
			foundErrors = append(foundErrors, foundError)
		}

		if err := input.ValidateTitleBaseSchemaType(note.Name()); err != nil {
			foundError := fmt.Errorf(
				"invalid base schema type found in '%s' file: %w", note.Name(), err,
			)
			foundErrors = append(foundErrors, foundError)
		}
	}

	log.Info().
		Int("count", len(foundErrors)).
		Msgf("found %d validation errors", len(foundErrors))

	for _, foundError := range foundErrors {
		log.Warn().Msg(foundError.Error())
	}

	return nil
}
