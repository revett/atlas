package cmd

import (
	"fmt"
	"os/exec"

	"github.com/revett/sepia/internal/note"
	"github.com/revett/sepia/internal/schema"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	autoDoctor  bool
	codeSnippet bool
)

// Root returns a cobra.Command type that acts as the entrypoint CLI command.
func Root() *cobra.Command {
	root := cobra.Command{
		Use:       "sepia {area|entity|meeting|project|review|scratch|system}",
		Example:   "sepia scratch",
		Short:     "CLI focused personal knowledge management tool",
		ValidArgs: schema.Schemas(),
		RunE:      rootRunE,
		Args: cobra.MatchAll(
			cobra.ExactArgs(1), cobra.OnlyValidArgs,
		),
	}

	root.Flags().BoolVarP(
		&autoDoctor,
		"auto-doctor",
		"a",
		true,
		"Run the doctor command before creating a new note",
	)

	root.Flags().BoolVarP(
		&codeSnippet,
		"code-snippet",
		"c",
		false,
		"Append a code block to the bottom of the new note",
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

	if codeSnippet {
		log.Info().Msg("--code-snippet flag enabled")
	}

	n, err := note.NewNote(schema)
	if err != nil {
		return fmt.Errorf("failed to create new note type: %w", err)
	}

	filepath, err := n.WriteToDisk(codeSnippet)
	if err != nil {
		return fmt.Errorf("failed to create new note: %w", err)
	}

	if err := exec.Command("code", filepath).Run(); err != nil { //nolint:gosec
		return fmt.Errorf("failed to open new note in vscode: %w", err)
	}

	return nil
}
