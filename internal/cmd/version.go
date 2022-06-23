package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// Version returns a cobra.Command type that outputs the version of the CLI.
func Version(version string) *cobra.Command {
	return &cobra.Command{
		Use:                   "version",
		Short:                 "Print the version",
		DisableFlagsInUseLine: true,
		Args:                  cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			fmt.Println( // nolint:forbidigo
				strings.TrimSpace(version),
			)
		},
	}
}
