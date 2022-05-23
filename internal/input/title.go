package input

import (
	"fmt"
	"regexp"
)

const exp = `^([a-z0-9]+(?:-[a-z0-9]+)*\.)+([a-z0-9]+(?:-[a-z0-9]+)*)$`

// ValidateTitle checks that a given string matches a combined kebab-case dot
// notation (e.g. area.foo.bar-123.baz).
func ValidateTitle(title string) error {
	ok, err := regexp.MatchString(exp, title)
	if err != nil {
		return fmt.Errorf("failed when matching string with regex: %w", err)
	}

	if !ok {
		return fmt.Errorf("title does not match kebab-case dot notation: %s", title)
	}

	return nil
}
