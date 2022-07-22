package cmd

import (
	"fmt"
	"os/exec"

	"github.com/revett/sepias/internal/note"
	"github.com/revett/sepias/internal/note/hierarchy"
	"github.com/spf13/cobra"
)

// Root returns a cobra.Command type that acts as the entrypoint CLI command.
func Root() *cobra.Command {
	return &cobra.Command{
		Use:                   "sepias {area|entity|interview|project|review|scratch|system}",
		Example:               "sepias scratch",
		DisableFlagsInUseLine: true,
		Short:                 "Tool that @revett uses to manage his notes",
		Args:                  cobra.ExactValidArgs(1),
		ValidArgs:             hierarchy.Schemas(),
		RunE:                  rootRunE,
	}
}

func rootRunE(c *cobra.Command, args []string) error {
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
