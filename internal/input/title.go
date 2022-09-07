package input

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/revett/sepia/internal/schema"
)

const exp = `^([a-z0-9]+(?:-[a-z0-9]+)*\.)+([a-z0-9]+(?:-[a-z0-9]+)*)$`

// ValidateTitleFormat checks that a given string matches a combined kebab-case
// dot notation (e.g. area.foo.bar-123.baz).
func ValidateTitleFormat(t string) error {
	withoutExtension := strings.TrimSuffix(
		t, filepath.Ext(t),
	)

	ok, err := regexp.MatchString(exp, withoutExtension)
	if err != nil {
		return ErrFailedRegexMatch
	}

	if !ok {
		return ErrInvalidTitleFormat
	}

	return nil
}

// ValidateTitleBaseSchemaType checks that the base schema type for a given note
// title (e.g. "area" from "area.language.go.error") is valid.
func ValidateTitleBaseSchemaType(t string) error {
	parts := strings.Split(t, ".")

	if len(parts) <= 1 {
		return ErrInsufficientNumberOfTitlePartsErr
	}

	schemas := []string{
		schema.ArchiveSchema,
	}
	schemas = append(schemas, schema.Schemas()...)

	for _, s := range schemas {
		if parts[0] == s {
			return nil
		}
	}

	return ErrUnrecognisedBaseSchemaTypeErr
}
