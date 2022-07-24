package note

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/revett/sepias/internal/input"
	"github.com/revett/sepias/internal/note/hierarchy"
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
	for _, e := range hierarchy.Schemas() {
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

	title, err := generateNoteTitle(note)
	if err != nil {
		return Note{}, fmt.Errorf("failed to generate title for new note: %w", err)
	}

	title = fmt.Sprintf("%s.%s", schema, title)

	err = input.ValidateTitleFormat(title)
	if err != nil {
		return Note{}, fmt.Errorf("invalid title format: %w", err)
	}
	note.title = title

	return note, nil
}

// CreateNote checks that the new note does not already exist, then creates the
// new note file, and appends contents to the file (header, template).
func CreateNote(note Note) (string, error) {
	filepath := fmt.Sprintf("./%s.md", note.title)

	if _, err := os.Stat(filepath); err == nil {
		return "", fmt.Errorf("note already exists: %s", filepath)
	}

	header, err := generateFrontmatterHeader()
	if err != nil {
		return "", err
	}

	templatePath, err := findTemplate(note.title)
	if err != nil {
		return "", err
	}

	tmpl, err := readTemplate(templatePath)
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

func generateNoteTitle(note Note) (string, error) { // nolint:funlen,cyclop
	var fn func() (string, error)

	switch note.schema {
	case hierarchy.AreaSchema:
		fn = func() (string, error) {
			return readInput(
				hierarchy.AreaSchema,
				[]string{
					"language.go.errors",
				},
			)
		}
	case hierarchy.EntitySchema:
		fn = func() (string, error) {
			return readInput(
				hierarchy.EntitySchema,
				[]string{
					"person.colleague.john-smith",
				},
			)
		}
	case hierarchy.MeetingSchema:
		fn = func() (string, error) {
			input, err := readInput(
				hierarchy.MeetingSchema,
				[]string{
					"design.2022-q3-review",
					"interview.cultural",
					"interview.technical",
				},
			)
			if err != nil {
				return "", err
			}

			if input == "interview.cultural" || input == "interview.technical" {
				return fmt.Sprintf(
					"%s.%s", input, time.Now().Format("2006.01.02.1504"),
				), nil
			}

			return input, nil
		}
	case hierarchy.ProjectSchema:
		fn = func() (string, error) {
			return readInput(
				hierarchy.ProjectSchema,
				[]string{
					"video-app.mvp-features",
				},
			)
		}
	case hierarchy.ReviewSchema:
		fn = func() (string, error) {
			y, w := time.Now().ISOWeek()
			return fmt.Sprintf("%d.%d", y, w), nil
		}
	case hierarchy.ScratchSchema:
		fn = func() (string, error) {
			return time.Now().Format("2006.01.02.150405"), nil
		}
	case hierarchy.SystemSchema:
		fn = func() (string, error) {
			return readInput(
				hierarchy.SystemSchema,
				[]string{
					"monthly-accounts",
				},
			)
		}
	}

	return fn()
}

func readInput(schema string, examples []string) (string, error) {
	p := tea.NewProgram(
		input.NewModel(schema, examples),
	)

	m, err := p.StartReturningModel()
	if err != nil {
		return "", fmt.Errorf("failed when starting tui model: %w", err)
	}

	return m.View(), nil
}
