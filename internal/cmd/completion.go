package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Completion returns a cobra.Command type that generates shell completion
// scripts for different types of environments.
func Completion() *cobra.Command {
	return &cobra.Command{
		Use:                   "completion [bash|zsh|fish|powershell]",
		Short:                 "Generate shell completion script",
		DisableFlagsInUseLine: true,
		Args:                  cobra.ExactValidArgs(1),
		ValidArgs:             []string{"bash", "fish", "powershell", "zsh"},
		RunE:                  completionRunE,
	}
}

func completionRunE(c *cobra.Command, args []string) error {
	switch args[0] {
	case "bash":
		return Root().GenBashCompletion(os.Stdout)
	case "fish":
		return Root().GenFishCompletion(os.Stdout, true)
	case "powershell":
		return Root().GenPowerShellCompletionWithDesc(os.Stdout)
	case "zsh":
		return Root().GenZshCompletion(os.Stdout)
	default:
		return nil
	}
}
