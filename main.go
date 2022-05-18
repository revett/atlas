package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/revett/sepias/internal/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const phraseChunkLength = 4

func main() {
	rand.Seed(time.Now().UnixNano())

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out: os.Stdout,
		},
	)

	if err := cmd.Root().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// 	argsWithoutProg := os.Args[1:]
// 	if len(argsWithoutProg) != 1 {
// 		log.Fatal().Msg("must pass exactly one argument to cli")
// 	}
// 	noteType := strings.ToLower(argsWithoutProg[0])

// 	if ok := validateNoteType(noteType); !ok {
// 		log.Fatal().Str("noteType", noteType).Msg("invalid note type")
// 	}

// 	var title string
// 	var err error

// 	// TODO: add option to disable colors.

// 	switch noteType {
// 	case system.String():
// 		q := fmt.Sprintf(
// 			"%s %s",
// 			aurora.Magenta(system.String()).Bold(),
// 			aurora.Cyan("(e.g. monthly-accounts)"),
// 		)
// 		title, err = readInput(q, system.String())
// 	case project.String():
// 		q := fmt.Sprintf(
// 			"%s %s",
// 			aurora.Magenta(project.String()).Bold(),
// 			aurora.Cyan("(e.g. video-app.mvp-features)"),
// 		)
// 		title, err = readInput(q, project.String())
// 	case entity.String():
// 		q := fmt.Sprintf(
// 			"%s %s",
// 			aurora.Magenta(entity.String()).Bold(),
// 			aurora.Cyan("(e.g. colleague.john-smith)"),
// 		)
// 		title, err = readInput(q, entity.String())
// 	case interview.String():
// 		title = time.Now().Format("2006.01.02.1504")
// 	case area.String():
// 		q := fmt.Sprintf(
// 			"%s %s",
// 			aurora.Magenta(area.String()).Bold(),
// 			aurora.Cyan("(e.g. language.go.errors)"),
// 		)
// 		title, err = readInput(q, area.String())
// 	case scratch.String():
// 		title = time.Now().Format("2006.01.02.150405")
// 	}
// 	if err != nil {
// 		log.Fatal().Err(err).Send()
// 	}

// 	// TODO: validate title.

// 	filepath, err := createNote(noteType, title)
// 	if err != nil {
// 		log.Fatal().Err(err).Send()
// 	}

// 	err = exec.Command("code", filepath).Run() // nolint:gosec
// 	if err != nil {
// 		log.Fatal().Err(err).Send()
// 	}
// }

// const (
// 	system noteType = iota
// 	project
// 	entity
// 	interview
// 	area
// 	scratch
// )

// type noteType int

// // String implements the Stringer.String() interface.
// func (n noteType) String() string {
// 	switch n {
// 	case system:
// 		return "system"
// 	case project:
// 		return "project"
// 	case entity:
// 		return "entity"
// 	case interview:
// 		return "interview"
// 	case area:
// 		return "area"
// 	case scratch:
// 		return "scratch"
// 	default:
// 		return "unknown"
// 	}
// }

// func createNote(noteType string, title string) (string, error) {
// 	filepath := fmt.Sprintf("./%s.%s.md", noteType, title)

// 	if _, err := os.Stat(filepath); err == nil {
// 		return "", fmt.Errorf("note already exists: %s", filepath)
// 	}

// 	header, err := generateFrontmatter()
// 	if err != nil {
// 		return "", err
// 	}

// 	tmpl, err := readTemplate(noteType)
// 	if err != nil {
// 		return "", err
// 	}

// 	file, err := os.Create(filepath) // nolint:gosec
// 	if err != nil {
// 		return "", fmt.Errorf("unable to create new note file: %w", err)
// 	}
// 	defer func() {
// 		if err := file.Close(); err != nil {
// 			log.Warn().Err(err).Msg("error occurred when closing file during defer")
// 		}
// 	}()

// 	c := fmt.Sprintf("%s\n%s", header, tmpl)

// 	_, err = file.WriteString(c)
// 	if err != nil {
// 		return "", fmt.Errorf("unable to write template to new note file: %w", err)
// 	}

// 	return filepath, nil
// }

// // TODO: add option to disable cspell.
// func generateFrontmatter() (string, error) {
// 	format := `---
// // cspell:disable-next-line
// id: %s
// created: %d (%s)
// ---
// `

// 	id, err := language.RandomPhrase(phraseChunkLength)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to generate random phrase: %w", err)
// 	}

// 	return fmt.Sprintf(
// 		format, id, time.Now().Unix(), time.Now().Format(time.RFC1123),
// 	), nil
// }

// func readInput(question string, prefix string) (string, error) {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Println(question) // nolint:forbidigo
// 	s := fmt.Sprintf("> %s.", prefix)
// 	fmt.Print(aurora.Faint(s)) // nolint:forbidigo

// 	answer, err := reader.ReadString('\n')
// 	if err != nil {
// 		return "", fmt.Errorf("failed to read input from use: %w", err)
// 	}

// 	answer = strings.ReplaceAll(answer, "\n", "")
// 	return answer, nil
// }

// func readTemplate(noteType string) (string, error) {
// 	filepath := fmt.Sprintf("./templates/%s.md", noteType)

// 	f, err := os.Open(filepath) // nolint:gosec
// 	if err != nil {
// 		return "", fmt.Errorf("failed to open template file: %w", err)
// 	}

// 	b, err := io.ReadAll(f)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to read contents of template file: %w", err)
// 	}

// 	return string(b), nil
// }

// func validateNoteType(input string) bool {
// 	types := []noteType{
// 		system,
// 		project,
// 		entity,
// 		interview,
// 		area,
// 		scratch,
// 	}

// 	for _, e := range types {
// 		if e.String() == input {
// 			return true
// 		}
// 	}

// 	return false
// }
