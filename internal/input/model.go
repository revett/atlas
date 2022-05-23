package input

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model is a custom github.com/charmbracelet/bubbletea#Model that reads the
// title of a new note.
type Model struct {
	complete bool
	example  string
	input    textinput.Model
	schema   string
}

const (
	characterLimit = 128
	inputWidth     = 64
)

var (
	exampleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("44"))
	schemaStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
)

// NewModel creates a new Model type.
func NewModel(schema string, example string) Model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = characterLimit
	ti.Width = inputWidth
	ti.Prompt = fmt.Sprintf("> %s.", schema)

	return Model{
		input:   ti,
		example: example,
		schema:  schema,
	}
}

// Init implements the bubbletea.Model.Init() interface.
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update implements the bubbletea.Model.Update() interface.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			m.complete = true
			return m, tea.Quit
		case tea.KeyTab:
			// TODO: implement tab completion using existing note hierachy.
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

// View implements the bubbletea.Model.View() interface.
func (m Model) View() string {
	// TODO: fix that this causes the question input to disappear after enter.
	if m.complete {
		return m.input.Value()
	}

	return fmt.Sprintf(
		"%s (e.g. %s)\n%s\n",
		schemaStyle.Render(m.schema),
		exampleStyle.Render(m.example),
		m.input.View(),
	)
}
