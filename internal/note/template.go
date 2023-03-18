package note

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/revett/atlas/internal/config"
)

func appendCodeSnippet(c string) string {
	s := "## Code Snippet\n\n" + "```\n\n" + "```"
	return fmt.Sprintf("%s\n%s", c, s)
}

func findTemplate(cfg config.Config, title string) (string, error) {
	templatesDirectoryPath := filepath.ToSlash(
		path.Join(cfg.Path, "templates"),
	)

	files, err := os.ReadDir(templatesDirectoryPath)
	if err != nil {
		return "", ErrFindTemplatesReadDir
	}

	templates := []string{}

	for _, file := range files {
		if !file.Type().IsRegular() {
			continue
		}

		if filepath.Ext(file.Name()) != ".md" {
			continue
		}

		withoutExtension := strings.TrimSuffix(
			file.Name(), filepath.Ext(file.Name()),
		)

		templates = append(templates, withoutExtension)
	}

	sort.Slice(
		templates,
		func(i, j int) bool {
			x := len(strings.Split(templates[i], "."))
			y := len(strings.Split(templates[j], "."))
			return x > y
		},
	)

	for _, t := range templates {
		if strings.HasPrefix(title, t) {
			templateFilename := fmt.Sprintf("%s.md", t)
			return filepath.ToSlash(
				path.Join(templatesDirectoryPath, templateFilename),
			), nil
		}
	}

	return "", ErrMissingTemplates
}

func readTemplate(p string) (string, error) {
	f, err := os.Open(p) //nolint:gosec
	if err != nil {
		return "", fmt.Errorf("opening template file: %w", err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("reading contents of template file: %w", err)
	}

	return string(b), nil
}
