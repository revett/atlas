package cmd

import (
	"fmt"

	"github.com/revett/sepias/internal/note"
	"github.com/spf13/cobra"
)

// Root returns a cobra.Command type that acts as the entrypoint CLI command.
func Root() *cobra.Command {
	return &cobra.Command{
		Use:       "sepias",
		Short:     "Tool that @revett uses to manage his notes",
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: note.Types(),
		RunE:      runE,
	}
}

func runE(c *cobra.Command, args []string) error {
	fmt.Println("foo")
	return nil
}
