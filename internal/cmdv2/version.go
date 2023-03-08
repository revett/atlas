package cmdv2

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/revett/atlas/internal/config"
	"github.com/spf13/cobra"
)

func init() {
	RegisterCommand(Version{})
}

type Version struct {
	config config.Config
}

func (v Version) Command() *cobra.Command {
	return &cobra.Command{
		Use:                   "version",
		Short:                 "Print the version",
		DisableFlagsInUseLine: true,
		Args:                  cobra.NoArgs,
		RunE:                  v.runE,
	}
}

func (v Version) Init() tea.Cmd {
	return nil
}

func (v Version) Type() CommandType {
	return CommandTypeVersion
}

func (v Version) Update(msg tea.Msg) (tea.Model, tea.Cmd) { //nolint:ireturn
	return v, tea.Quit
}

func (v Version) View() string {
	return v.config.Version
}

func (v *Version) runE(c *cobra.Command, args []string) error {
	cfg, ok := c.Context().Value(config.ContextConfigKey).(config.Config)
	if !ok {
		return config.ErrContextConfigValueIsNotConfigType
	}

	v.config = cfg

	if _, err := tea.NewProgram(v).Run(); err != nil {
		return fmt.Errorf("running tui program: %w", err)
	}

	return nil
}
