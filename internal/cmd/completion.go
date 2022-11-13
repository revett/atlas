package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Completion returns a cobra.Command type that generates shell completion
// scripts for different types of environments.
func Completion() *cobra.Command {
	return &cobra.Command{
		Use:                   "completion {bash|fish|powershell|zsh}",
		Short:                 "Generate shell completion script",
		Long:                  completionLong(Root().Name()),
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "fish", "powershell", "zsh"},
		RunE:                  completionRunE,
		Args: cobra.MatchAll(
			cobra.ExactArgs(1), cobra.OnlyValidArgs,
		),
	}
}

func completionRunE(c *cobra.Command, args []string) error {
	switch args[0] {
	case "bash":
		if err := Root().GenBashCompletion(os.Stdout); err != nil {
			return fmt.Errorf("failed to generate bash completion: %w", err)
		}
	case "fish":
		if err := Root().GenFishCompletion(os.Stdout, true); err != nil {
			return fmt.Errorf("failed to generate fish completion: %w", err)
		}
	case "powershell":
		if err := Root().GenPowerShellCompletionWithDesc(os.Stdout); err != nil {
			return fmt.Errorf("failed to generate powershell completion: %w", err)
		}
	case "zsh":
		if err := Root().GenZshCompletion(os.Stdout); err != nil {
			return fmt.Errorf("failed to generate zsh completion: %w", err)
		}
	}

	return nil
}

func completionLong(rootName string) string {
	return fmt.Sprintf(`Generate shell completion script.

Bash:
	$ source <(%[1]s completion bash)

	# To load completions for each session, execute once:
	# Linux:
	$ %[1]s completion bash > /etc/bash_completion.d/%[1]s
	# macOS:
	$ %[1]s completion bash > /usr/local/etc/bash_completion.d/%[1]s

Zsh:
	# If shell completion is not already enabled in your environment,
	# you will need to enable it.  You can execute the following once:

	$ echo "autoload -U compinit; compinit" >> ~/.zshrc

	# To load completions for each session, execute once:
	$ sudo mkdir -p ${fpath[1]}
	$ %[1]s completion zsh > _%[1]s
	$ mv _%[1]s ${fpath[1]}/_%[1]s

	# You will need to start a new shell for this setup to take effect.

fish:
	$ %[1]s completion fish | source

	# To load completions for each session, execute once:
	$ %[1]s completion fish > ~/.config/fish/completions/%[1]s.fish

PowerShell:
	PS> %[1]s completion powershell | Out-String | Invoke-Expression

	# To load completions for every new session, run:
	PS> %[1]s completion powershell > %[1]s.ps1
	# and source this file from your PowerShell profile.
	`,
		rootName,
	)
}
