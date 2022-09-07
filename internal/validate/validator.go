package validate

// Validator is an interface for all validator types to derive from, it allows
// validation logic to be separate from the model.
type Validator interface {
	Validate(filename string, filenameWithoutExtension string) []error
}
