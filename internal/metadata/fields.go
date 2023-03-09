package metadata

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	minimumIDParts  = 3
	idCommentSuffix = " // cspell:disable-line"
)

// Fields are the different values within the Front Matter block at the top of any note.
type Fields struct {
	ID      string `yaml:"id"`
	Created string `yaml:"created"`
}

// Validate checks that the values within the Front Matter metadata block valid.
func (f Fields) Validate() error {
	if !strings.HasSuffix(f.ID, idCommentSuffix) {
		return fmt.Errorf("id field '%s' does not have required suffix: '%s'", f.ID, idCommentSuffix)
	}

	idWithoutComment := strings.Replace(f.ID, idCommentSuffix, "", 1)

	if parts := strings.Split(idWithoutComment, "-"); len(parts) < minimumIDParts {
		return fmt.Errorf("id field '%s' must have at least %d parts", idWithoutComment, minimumIDParts)
	}

	kebabCasePattern := regexp.MustCompile(`^[a-z]+(-[a-z]+)*$`)
	if !kebabCasePattern.MatchString(idWithoutComment) {
		return fmt.Errorf("id field '%s' must use kebab case format", idWithoutComment)
	}

	return nil
}
