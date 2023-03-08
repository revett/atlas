package validate

import (
	"regexp"
	"strings"

	"github.com/revett/atlas/internal/schema"
)

const (
	filenameFormatExp            = `^([a-z0-9]+(?:-[a-z0-9]+)*\.)+([a-z0-9]+(?:-[a-z0-9]+)*)\.md$`
	minimumNumberOfFilenameParts = 2
)

type filenameValidator struct{}

// NewFilenameValidator returns a filenameValidator type.
func NewFilenameValidator() filenameValidator {
	return filenameValidator{}
}

// Validate implements the validate.Validator.Valid() interface. It carries out
// a number of checks:
//   - If the file extension is correct
//   - Has a valid base schema (e.g. "area" from "area.language.go.error")
//   - If the filename matches a combined kebab-case dot notation
//     (e.g. "area.languagego.error-handling.md")
//   - If the filename has the required number of parts (e.g. not "area.md").
func (f filenameValidator) Validate(filename string) []error {
	var errs []error

	if err := validateFilenameFileExtension(filename); err != nil {
		errs = append(errs, err)
	}

	if err := validateFilenameBaseSchemaType(filename); err != nil {
		errs = append(errs, err)
	}

	if err := validateFilenameFormat(filename); err != nil {
		errs = append(errs, err)
	}

	if err := validateFilenameNumberOfParts(filename); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func validateFilenameBaseSchemaType(filename string) error {
	schemas := []string{
		schema.ArchiveSchema,
	}
	schemas = append(schemas, schema.Schemas()...)

	parts := strings.Split(filename, ".")

	for _, s := range schemas {
		if parts[0] == s {
			return nil
		}
	}

	return ErrUnrecognisedBaseSchemaType
}

func validateFilenameFileExtension(filename string) error {
	if !strings.HasSuffix(filename, ".md") {
		return ErrMissingMarkdownFileExtension
	}

	return nil
}

func validateFilenameFormat(filename string) error {
	ok, err := regexp.MatchString(filenameFormatExp, filename)
	if err != nil {
		return ErrFailedRegexMatch
	}

	if !ok {
		return ErrInvalidFilenameFormat
	}

	return nil
}

func validateFilenameNumberOfParts(filename string) error {
	parts := strings.Split(filename, ".")
	if len(parts) <= minimumNumberOfFilenameParts {
		return ErrInsufficientNumberOfFilenameParts
	}

	return nil
}
