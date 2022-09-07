package input

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rs/zerolog/log"
)

type (
	// Model is a custom github.com/charmbracelet/bubbletea#Model that reads the
	// filename of a new note.
	Model struct {
		complete bool
		examples []string
		input    textinput.Model
		schema   string
	}

	// ReturnValue is a hack to allow the charmbracelet/bubbletea#Model to return
	// the latest value from the terminal input whilst also preserving the prompt.
	// This is due to charmbracelet/bubbletea#Program.StartReturningModel
	// returning an interface which does not expose this value, thus instead when
	// the .Update() function of model has completed, the .View() function of the
	// model will return a JSON string instead of the raw prompt output. This
	// allows the consuming package to print the prompt a final time, thus
	// preserving it for the user, as well as having access to the raw value.
	ReturnValue struct {
		Filename string `json:"filename"`
		Prompt   string `json:"prompt"`
	}
)

const (
	characterLimit = 128
	inputWidth     = 64
)

// NewModel creates a new Model type.
func NewModel(schema string, examples []string) Model {
	input := textinput.New()
	input.Focus()
	input.CharLimit = characterLimit
	input.Width = inputWidth
	input.Prompt = fmt.Sprintf("> %s.", schema)

	return Model{
		input:    input,
		examples: examples,
		schema:   schema,
	}
}

// Init implements the bubbletea.Model.Init() interface.
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update implements the bubbletea.Model.Update() interface.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { //nolint:ireturn
	var cmd tea.Cmd

	switch msg := msg.(type) { //nolint:gocritic
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			m.complete = true
			return m, tea.Quit
		case tea.KeyTab:
			// TODO: implement tab completion using existing note hierarchy.
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

// View implements the bubbletea.Model.View() interface.
func (m Model) View() string {
	exampleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("44"))
	schemaStyle := lipgloss.NewStyle().Foreground(
		lipgloss.Color("205"),
	).Bold(true)

	renderedExamples := []string{}

	for _, e := range m.examples {
		renderedExamples = append(
			renderedExamples,
			exampleStyle.Render(e),
		)
	}

	prompt := fmt.Sprintf(
		"%s (e.g. %s)\n%s\n",
		schemaStyle.Render(m.schema),
		strings.Join(renderedExamples, ", "),
		m.input.View(),
	)

	if m.complete {
		rv := ReturnValue{
			Filename: m.input.Value(),
			Prompt:   prompt,
		}

		bytes, err := json.Marshal(&rv)
		if err != nil {
			log.Error().Err(err).
				Msg("failed when json marshalling filename input return value")
			return ""
		}

		return string(bytes)
	}

	return prompt
}
