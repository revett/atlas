package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Root returns a cobra.Command type that acts as the entrypoint CLI command.
func Root() *cobra.Command {
	return &cobra.Command{
		Use:   "sepias",
		Short: "Tool that @revett uses to manage his notes",
		RunE:  runE,
	}
}

func runE(c *cobra.Command, args []string) error {
	fmt.Println("foo")
	return nil
}
