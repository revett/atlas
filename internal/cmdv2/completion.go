package cmdv2

import (
	"bytes"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/revett/atlas/internal/config"
	"github.com/spf13/cobra"
)

func init() {
	RegisterCommand(Completion{})
}

// Completion is a CLI command that generates shell completion scripts for different types of
// environments.
type Completion struct {
	cmd  *cobra.Command
	args []string
}

func (c Completion) Command() *cobra.Command {
	return &cobra.Command{
		Use:                   "completion {bash|fish|powershell|zsh}",
		Short:                 "Generate shell completion script",
		Long:                  c.longCommandDescription(),
		DisableFlagsInUseLine: true,
		RunE:                  c.runE,
		Args: cobra.MatchAll(
			cobra.ExactArgs(1), cobra.OnlyValidArgs,
		),
		ValidArgs: []string{
			"bash",
			"fish",
			"powershell",
			"zsh",
		},
	}
}

func (c Completion) Init() tea.Cmd {
	return nil
}

func (c Completion) Type() CommandType {
	return CommandTypeCompletion
}

func (c Completion) Update(msg tea.Msg) (tea.Model, tea.Cmd) { //nolint:ireturn
	return c, tea.Quit
}

func (c Completion) View() string {
	buf := new(bytes.Buffer)

	switch c.args[0] {
	case "bash":
		if err := c.cmd.Root().GenBashCompletion(buf); err != nil {
			err := fmt.Errorf("generating bash completion: %w", err)
			return outputError(err)
		}
	case "fish":
		if err := c.cmd.Root().GenFishCompletion(buf, true); err != nil {
			err := fmt.Errorf("generating fish completion: %w", err)
			return outputError(err)
		}
	case "powershell":
		if err := c.cmd.Root().GenPowerShellCompletionWithDesc(buf); err != nil {
			err := fmt.Errorf("generating powershell completion: %w", err)
			return outputError(err)
		}
	case "zsh":
		if err := c.cmd.Root().GenZshCompletion(buf); err != nil {
			err := fmt.Errorf("generating zsh completion: %w", err)
			return outputError(err)
		}
	}

	return buf.String()
}

func (c *Completion) runE(cmd *cobra.Command, args []string) error {
	c.cmd = cmd
	c.args = args

	if _, err := tea.NewProgram(c).Run(); err != nil {
		return fmt.Errorf("running completion cli command: %w", err)
	}

	return nil
}

func (c Completion) longCommandDescription() string {
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
		config.CLIName,
	)
}
