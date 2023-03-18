package note

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/revett/atlas/internal/config"
	"github.com/revett/atlas/internal/file"
	"github.com/revett/atlas/internal/input"
	"github.com/revett/atlas/internal/metadata"
	"github.com/revett/atlas/internal/schema"
	"github.com/revett/atlas/internal/validate"
	"github.com/rs/zerolog/log"
)

// Note holds configuration for a new note, as well as methods for creating it.
type Note struct {
	Filename string
}

// NewNote creates a new Note type, whilst also validating the schema argument
// against known valid schemas and generates a filename for the new note.
func NewNote(noteSchema string) (Note, error) {
	// TODO: refactor to remove notion of title vs filepath

	valid := false
	for _, e := range schema.Schemas() {
		if e == noteSchema {
			valid = true
		}
	}

	if !valid {
		return Note{}, fmt.Errorf("unknown schema: %s", noteSchema)
	}

	filename, err := generateNoteFilename(noteSchema)
	if err != nil {
		return Note{}, fmt.Errorf(
			"generating filename for new note: %w", err,
		)
	}

	log.Info().Str("filename", filename).Msg("new note")

	note := Note{
		Filename: filename,
	}

	if errs := validate.NewFilenameValidator().Validate(note.Filename); errs != nil {
		return Note{}, fmt.Errorf("note filename validation failed: %v", errs)
	}

	return note, nil
}

// WriteToDisk checks that the new note does not already exist, then creates the
// new note file, and appends contents to the file (header, template).
func (n Note) WriteToDisk(cfg config.Config, codeSnippet bool) (string, error) {
	notePath := filepath.ToSlash(
		path.Join(cfg.Path, n.Filename),
	)

	log.Info().Str("path", notePath).Msg("writing to knowledge base")

	if err := file.DirectoryOrFileExists(notePath); err == nil {
		log.Warn().Str("path", notePath).Msg("note already exists")
		log.Info().Msg("opening note")
		return notePath, nil
	}

	templatePath, err := findTemplate(cfg, n.Filename)
	if err != nil {
		return "", err
	}

	tmpl, err := readTemplate(templatePath)
	if err != nil {
		return "", err
	}

	file, err := os.Create(notePath) //nolint:gosec
	if err != nil {
		return "", fmt.Errorf("unable to create new note file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Warn().Err(err).Msg("error occurred when closing file during defer")
		}
	}()

	content := fmt.Sprintf("%s\n%s", metadata.Generate(), tmpl)
	if codeSnippet {
		content = appendCodeSnippet(content)
	}

	if _, err = file.WriteString(content); err != nil {
		return "", fmt.Errorf("unable to write template to new note file: %w", err)
	}

	return notePath, nil
}

func generateNoteFilename(noteSchema string) (string, error) { //nolint:funlen
	var genFunc func() (string, error)

	switch noteSchema {
	case schema.AreaSchema:
		genFunc = func() (string, error) {
			return readInput(
				schema.AreaSchema,
				[]string{
					"language.go.errors",
				},
			)
		}

	case schema.EntitySchema:
		genFunc = func() (string, error) {
			return readInput(
				schema.EntitySchema,
				[]string{
					"person.colleague.john-smith",
				},
			)
		}

	case schema.MeetingSchema:
		genFunc = func() (string, error) {
			now := time.Now().Format("2006.01.02.1504")

			input, err := readInput(
				fmt.Sprintf("%s.%s", schema.MeetingSchema, now),
				[]string{
					"design.2022-q3-review",
					"interview.cultural",
					"interview.technical",
				},
			)
			if err != nil {
				return "", err
			}

			return fmt.Sprintf("%s.%s", now, input), nil
		}

	case schema.ProjectSchema:
		genFunc = func() (string, error) {
			return readInput(
				schema.ProjectSchema,
				[]string{
					"video-app.mvp-features",
				},
			)
		}

	case schema.ScratchSchema:
		genFunc = func() (string, error) {
			return time.Now().Format("2006.01.02.150405"), nil
		}

	case schema.SystemSchema:
		genFunc = func() (string, error) {
			return readInput(
				schema.SystemSchema,
				[]string{
					"monthly-accounts",
				},
			)
		}
	}

	filename, err := genFunc()
	if err != nil {
		return "", fmt.Errorf("generating filename: %w", err)
	}

	return fmt.Sprintf("%s.%s.md", noteSchema, filename), nil
}

func readInput(schema string, examples []string) (string, error) {
	p := tea.NewProgram(
		input.NewModel(schema, examples),
	)

	model, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("starting tui model: %w", err)
	}

	var returnValue input.ReturnValue

	err = json.Unmarshal([]byte(model.View()), &returnValue)
	if err != nil {
		return "", fmt.Errorf(
			"json unmarshalling filename return value: %w", err,
		)
	}

	fmt.Printf(returnValue.Prompt) //nolint:forbidigo

	return returnValue.Filename, nil
}
