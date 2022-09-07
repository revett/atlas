package metadata

import (
	"fmt"
	"strings"
)

const (
	idFieldChunks        = 4
	idFieldCommentSuffix = " // cspell:disable-line"
)

// Fields are the different values within the Front Matter block at the top of
// any note.
type Fields struct {
	ID      string `yaml:"id"`
	Created string `yaml:"created"`
}

// Validate checks that the values within the Front Matter metadata block valid.
func (f Fields) Validate() error {
	if !strings.HasSuffix(f.ID, idFieldCommentSuffix) {
		return fmt.Errorf(
			"id field '%s' does not have required suffix: '%s'",
			f.ID,
			idFieldCommentSuffix,
		)
	}

	idWithoutComment := strings.Replace(f.ID, idFieldCommentSuffix, "", 1)

	if chunks := strings.Split(
		idWithoutComment, "-",
	); len(chunks) != idFieldChunks {
		return fmt.Errorf(
			"id field '%s' should have format: 'a-b-c-d'", idWithoutComment,
		)
	}

	return nil
}
