package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/revett/sepias/internal/note"
	"github.com/spf13/cobra"
)

// Root returns a cobra.Command type that acts as the entrypoint CLI command.
func Root() *cobra.Command {
	return &cobra.Command{
		Use:                   "sepias { system | entity | project | interview | area | scratch }",
		Example:               "sepias area",
		DisableFlagsInUseLine: true,
		Short:                 "Tool that @revett uses to manage his notes",
		Args:                  cobra.ExactValidArgs(1),
		ValidArgs:             note.Schemas(),
		RunE:                  runE,
	}
}

func runE(c *cobra.Command, args []string) error {
	note, err := note.NewNote(
		strings.ToLower(args[0]),
	)
	if err != nil {
		return fmt.Errorf("failed to create new note type: %w", err)
	}

	fp, err := note.Create()
	if err != nil {
		return fmt.Errorf("failed to create new note: %w", err)
	}

	err = exec.Command("code", fp).Run() // nolint:gosec
	if err != nil {
		return fmt.Errorf("failed to open new note in vscode: %w", err)
	}

	return nil
}