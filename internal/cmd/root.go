package cmd

import (
	"fmt"
	"os/exec"

	"github.com/revett/sepia/internal/note"
	"github.com/revett/sepia/internal/schema"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var autoDoctor bool //nolint:gochecknoglobals

// Root returns a cobra.Command type that acts as the entrypoint CLI command.
func Root() *cobra.Command {
	root := cobra.Command{
		Use:       "sepia {area|entity|meeting|project|review|scratch|system}",
		Example:   "sepia scratch",
		Short:     "CLI focused personal knowledge management tool",
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: schema.Schemas(),
		RunE:      rootRunE,
	}

	root.Flags().BoolVar(
		&autoDoctor,
		"auto-doctor",
		true,
		"Run the doctor command before creating a new note",
	)

	return &root
}

func rootRunE(c *cobra.Command, args []string) error {
	schema := args[0]

	if autoDoctor {
		log.Info().Msg("--auto-doctor flag enabled")

		// TODO: refactor to use underlying validator rather than command itself
		if err := Doctor().RunE(nil, nil); err != nil {
			return fmt.Errorf("failed to run the doctor command before: %w", err)
		}
	}

	n, err := note.NewNote(schema)
	if err != nil {
		return fmt.Errorf("failed to create new note type: %w", err)
	}

	filepath, err := n.WriteToDisk()
	if err != nil {
		return fmt.Errorf("failed to create new note: %w", err)
	}

	err = exec.Command("code", filepath).Run() //nolint:gosec
	if err != nil {
		return fmt.Errorf("failed to open new note in vscode: %w", err)
	}

	return nil
}
