package input

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/revett/sepias/internal/note/hierarchy"
)

const exp = `^([a-z0-9]+(?:-[a-z0-9]+)*\.)+([a-z0-9]+(?:-[a-z0-9]+)*)$`

// ValidateTitleFormat checks that a given string matches a combined kebab-case
// dot notation (e.g. area.foo.bar-123.baz).
func ValidateTitleFormat(title string) error {
	ok, err := regexp.MatchString(exp, title)
	if err != nil {
		return fmt.Errorf("failed when matching string with regex: %w", err)
	}

	if !ok {
		return fmt.Errorf("title does not match kebab-case dot notation: %s", title)
	}

	return nil
}

// ValidateTitleBaseSchemaType checks that the base schema type for a given note
// title (e.g. "area" from "area.language.go.error") is valid.
func ValidateTitleBaseSchemaType(t string) error {
	parts := strings.Split(t, ".")

	if len(parts) == 0 {
		return fmt.Errorf("title does not have sufficient number of parts")
	}

	schemas := []string{hierarchy.ArchiveSchema}
	schemas = append(schemas, hierarchy.Schemas()...)

	for _, s := range schemas {
		if parts[0] == s {
			return nil
		}
	}

	return fmt.Errorf("note uses an unrecognised base schema type: %s", parts[0])
}
