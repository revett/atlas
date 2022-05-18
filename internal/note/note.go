package note

import (
	"fmt"
	"os"
	"time"

	"github.com/logrusorgru/aurora/v3"
	"github.com/revett/sepias/internal/input"
	"github.com/rs/zerolog/log"
)

type Note struct {
	schema string
	title  string
}

func NewNote(s string) (Note, error) {
	valid := false
	for _, e := range Schemas() {
		if e == s {
			valid = true
		}
	}

	if !valid {
		return Note{}, fmt.Errorf("unknown schema: %s", s)
	}

	note := Note{
		schema: s,
	}

	t, err := note.generateTitle()
	if err != nil {
		return Note{}, fmt.Errorf("failed to generate title for new note: %w", err)
	}
	note.title = t

	return note, nil
}

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

	// TODO: validate title.

	return fn()
}

func readInput(schema string, example string) (string, error) {
	q := fmt.Sprintf(
		"%s (e.g. %s)", aurora.Magenta(schema).Bold(), aurora.Cyan(example),
	)

	t, err := input.Question(q, schema)
	if err != nil {
		return "", fmt.Errorf("failed when reading answer to question: %w", err)
	}

	return t, nil
}
