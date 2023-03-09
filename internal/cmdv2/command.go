package cmdv2

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

type (
	// Command is an interface that implements the bubbletea.Model interface, as well as returning a
	// cobra.Command type, so that they can play well together.
	Command interface {
		// Command returns the cobra.Command type so that the CLI can be built.
		Command() *cobra.Command

		// Init implements the bubbletea.Model.Init() interface.
		Init() tea.Cmd

		// Type returns a unique type string identifying the command.
		Type() CommandType

		// Update implements the bubbletea.Model.Update() interface.
		Update(msg tea.Msg) (tea.Model, tea.Cmd)

		// View implements the bubbletea.Model.View() interface.
		View() string
	}

	// CommandType is the unique ID for a Command.
	CommandType string
)

const (
	CommandTypeCompletion CommandType = "completion"
	CommandTypeVersion    CommandType = "version"
)
