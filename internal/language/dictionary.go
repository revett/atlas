package language

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Dictionary attempts to load the wordlist from '/usr/share/dict/words'; if it
// does not exist then it will return a nil slice.
func Dictionary() ([]string, error) {
	if _, err := os.Stat("/usr/share/dict/words"); errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}

	f, err := os.Open("/usr/share/dict/words")
	if err != nil {
		return nil, fmt.Errorf("failed to open '/usr/share/dict/words': %w", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read contents of '/usr/share/dict/words': %w", err,
		)
	}

	t := strings.Split(string(bytes), "\n")
	return t, nil
}
