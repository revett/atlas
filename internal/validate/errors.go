package validate

const (
	// ErrFailedRegexMatch is a sentinel error for when the check fails due to
	// regexp.MatchString returning an error, commonly caused by an invalid regex
	// pattern.
	ErrFailedRegexMatch Error = iota

	// ErrInsufficientNumberOfFilenameParts is a sentinel error for when the check
	//  fails due to the filename not having at least one part.
	ErrInsufficientNumberOfFilenameParts

	// ErrInvalidFilenameFormat is a sentinel error for when the check fails due
	// to a note filename using the incorrect format.
	ErrInvalidFilenameFormat

	// ErrMissingMarkdownFileExtension is a sentinel error for when the check
	// fails due to the filename not ending with the ".md" markdown file
	// extension.
	ErrMissingMarkdownFileExtension

	// ErrUnrecognisedBaseSchemaType is a sentinel error for when the check fails
	// due to the filename using a base schema type which is not valid/known.
	ErrUnrecognisedBaseSchemaType
)

// Error allows for sentinel errors via enums.
type Error uint8

// Error implements the errors.Error interface.
func (e Error) Error() string {
	switch e {
	case ErrFailedRegexMatch:
		return "matching string with regex"
	case ErrInsufficientNumberOfFilenameParts:
		return "filename does not have sufficient number of parts"
	case ErrInvalidFilenameFormat:
		return "filename does not match kebab-case dot notation"
	case ErrMissingMarkdownFileExtension:
		return "filename does not end in '.md' markdown file extension"
	case ErrUnrecognisedBaseSchemaType:
		return "note uses an unrecognised base schema type"
	default:
		return "unknown error"
	}
}
