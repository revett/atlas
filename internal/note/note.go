package note

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/revett/sepias/internal/input"
	"github.com/rs/zerolog/log"
)

// Note holds configuration for a new note, as well as methods for creating it.
type Note struct {
	schema string
	title  string
}

// NewNote creates a new Note type, whilst also validating the schema argument
// against known valid schemas and generates a title for the new note.
func NewNote(schema string) (Note, error) {
	valid := false
	for _, e := range Schemas() {
		if e == schema {
			valid = true
		}
	}

	if !valid {
		return Note{}, fmt.Errorf("unknown schema: %s", schema)
	}

	note := Note{
		schema: schema,
	}

	title, err := note.generateTitle()
	if err != nil {
		return Note{}, fmt.Errorf("failed to generate title for new note: %w", err)
	}

	err = input.ValidateTitle(title)
	if err != nil {
		return Note{}, fmt.Errorf("invalid title format: %w", err)
	}
	note.title = title

	return note, nil
}

// Create checks that the new note does not already exist, then creates the new
// note file, and appends contents to the file (header, template).
func (n Note) Create() (string, error) {
	filepath := fmt.Sprintf("./%s.%s.md", n.schema, n.title)

	if _, err := os.Stat(filepath); err == nil {
		return "", fmt.Errorf("note already exists: %s", filepath)
	}

	header, err := generateFrontmatterHeader()
	if err != nil {
		return "", err
	}

	tmpl, err := readSchemaTemplate(n.schema)
	if err != nil {
		return "", err
	}

	file, err := os.Create(filepath) // nolint:gosec
	if err != nil {
		return "", fmt.Errorf("unable to create new note file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Warn().Err(err).Msg("error occurred when closing file during defer")
		}
	}()

	c := fmt.Sprintf("%s\n%s", header, tmpl)

	_, err = file.WriteString(c)
	if err != nil {
		return "", fmt.Errorf("unable to write template to new note file: %w", err)
	}

	return filepath, nil
}

func (n Note) generateTitle() (string, error) {
	var fn func() (string, error)

	switch n.schema {
	case SystemSchema:
		fn = func() (string, error) {
			return readInput(SystemSchema, "monthly-accounts")
		}
	case ProjectSchema:
		fn = func() (string, error) {
			return readInput(ProjectSchema, "video-app.mvp-features")
		}
	case EntitySchema:
		fn = func() (string, error) {
			return readInput(EntitySchema, "colleague.john-smith")
		}
	case InterviewSchema:
		fn = func() (string, error) {
			return time.Now().Format("2006.01.02.1504"), nil
		}
	case AreaSchema:
		fn = func() (string, error) {
			return readInput(AreaSchema, "language.go.errors")
		}
	case ScratchSchema:
		fn = func() (string, error) {
			return time.Now().Format("2006.01.02.150405"), nil
		}
	}

	return fn()
}

func readInput(schema string, example string) (string, error) {
	p := tea.NewProgram(
		input.NewModel(schema, example),
	)

	m, err := p.StartReturningModel()
	if err != nil {
		return "", fmt.Errorf("failed when starting tui model: %w", err)
	}

	return m.View(), nil
}
