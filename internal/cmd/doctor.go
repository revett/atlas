package cmd

import (
	"fmt"

	"github.com/revett/atlas/internal/base"
	"github.com/revett/atlas/internal/config"
	"github.com/revett/atlas/internal/metadata"
	"github.com/revett/atlas/internal/validate"
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
	cfg, ok := c.Context().Value(config.ContextConfigKey).(config.Config)
	if !ok {
		return config.ErrContextConfigValueIsNotConfigType
	}

	log.Info().Str("path", cfg.Path).Msg("validating knowledge base")

	files, err := base.Read(cfg.Path)
	if err != nil {
		return fmt.Errorf("reading notes in knowledge base: %w", err)
	}

	log.Info().Int("count", len(files)).Msgf("found %d notes", len(files))

	var foundErrors []error

	for _, file := range files {
		filename := file.Name()

		if errs := validate.NewFilenameValidator().Validate(filename); errs != nil {
			foundError := fmt.Errorf(
				"./%s: invalid note filename: %v", filename, errs,
			)
			foundErrors = append(foundErrors, foundError)
		}

		metaFields, err := metadata.Parse(cfg, filename)
		if err != nil {
			foundError := fmt.Errorf(
				"./%s: unable to parse front matter metadata: %w", filename, err,
			)
			foundErrors = append(foundErrors, foundError)
		}

		if err := metaFields.Validate(); err != nil {
			foundError := fmt.Errorf(
				"./%s: invalid front matter metadata: %w", filename, err,
			)
			foundErrors = append(foundErrors, foundError)
		}
	}

	log.Info().
		Int("count", len(foundErrors)).
		Msgf("detected %d validation errors", len(foundErrors))

	for _, foundError := range foundErrors {
		log.Warn().Msg(foundError.Error())
	}

	return nil
}
