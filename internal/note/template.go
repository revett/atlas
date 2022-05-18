package note

import (
	"fmt"
	"io"
	"os"
)

func readSchemaTemplate(s string) (string, error) {
	filepath := fmt.Sprintf("./templates/%s.md", s)

	f, err := os.Open(filepath) // nolint:gosec
	if err != nil {
		return "", fmt.Errorf("failed to open template file: %w", err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("failed to read contents of template file: %w", err)
	}

	return string(b), nil
}
