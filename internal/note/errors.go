package note

const (
	// ErrFindTemplatesReadDir is a sentinel error for when the reading of the
	// templates directory fails.
	ErrFindTemplatesReadDir Error = iota

	// ErrMissingTemplates is a sentinel error for when no valid files are
	// found within the templates directory.
	ErrMissingTemplates
)

// Error allows for sentinel errors via enums.
type Error uint8

// Error implements the errors.Error interface.
func (e Error) Error() string {
	switch e {
	case ErrFindTemplatesReadDir:
		return "error when reading files within templates directory"
	case ErrMissingTemplates:
		return "no valid templates found locally"
	default:
		return "unknown error"
	}
}
