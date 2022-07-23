package cmd

import (
	"fmt"
	"os/exec"

	"github.com/revett/sepias/internal/note"
	"github.com/revett/sepias/internal/note/hierarchy"
	"github.com/spf13/cobra"
)

var autoDoctor bool

// Root returns a cobra.Command type that acts as the entrypoint CLI command.
func Root() *cobra.Command {
	root := cobra.Command{
		Use:       "sepias {area|entity|interview|project|review|scratch|system}",
		Example:   "sepias scratch",
		Short:     "Tool that @revett uses to manage his notes",
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: hierarchy.Schemas(),
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
	if autoDoctor {
		doctor := Doctor()
		if err := doctor.RunE(nil, nil); err != nil {
			return fmt.Errorf("failed to run the doctor command before: %w", err)
		}
	}

	n, err := note.NewNote(args[0])
	if err != nil {
		return fmt.Errorf("failed to create new note type: %w", err)
	}

	fp, err := note.CreateNote(n)
	if err != nil {
		return fmt.Errorf("failed to create new note: %w", err)
	}

	err = exec.Command("code", fp).Run() // nolint:gosec
	if err != nil {
		return fmt.Errorf("failed to open new note in vscode: %w", err)
	}

	return nil
}
