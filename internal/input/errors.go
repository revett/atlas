package input

const (
	// ErrFailedRegexMatch is a sentinel error for when the .ValidateTitleFormat
	// check fails due to regexp.MatchString returning an error, commonly caused
	// by an invalid regex pattern.
	ErrFailedRegexMatch Error = iota

	// ErrInsufficientNumberOfTitlePartsErr is a sentinel error for when the
	// .ValidateTitleBaseSchemaType check fails due to the title not having at
	// least one part.
	ErrInsufficientNumberOfTitlePartsErr

	// ErrInvalidTitleFormat is a sentinel error for when the .ValidateTitleFormat
	// check fails due to a note title using the incorrect format.
	ErrInvalidTitleFormat

	// ErrUnrecognisedBaseSchemaTypeErr is a sentinel error for when the
	// .ValidateTitleBaseSchemaType check fails due to the title using a base
	// schema type which is not valid/known.
	ErrUnrecognisedBaseSchemaTypeErr
)

// Error allows for sentinel errors via enums.
type Error uint8

// Error implements the errors.Error interface.
func (e Error) Error() string {
	switch e {
	case ErrFailedRegexMatch:
		return "failed when matching string with regex"
	case ErrInsufficientNumberOfTitlePartsErr:
		return "title does not have sufficient number of parts"
	case ErrInvalidTitleFormat:
		return "title does not match kebab-case dot notation"
	case ErrUnrecognisedBaseSchemaTypeErr:
		return "note uses an unrecognised base schema type"
	default:
		return "unknown error"
	}
}
