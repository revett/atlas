package schema

// TODO: refactor to be types with .Stringer interface, as well as filename
// generation func.

const (
	// ArchiveSchema represents the "archive" schema type.
	ArchiveSchema = "archive"

	// AreaSchema represents the "area" schema type.
	AreaSchema = "area"

	// EntitySchema represents the "entity" schema type.
	EntitySchema = "entity"

	// MeetingSchema represents the "meeting" schema type.
	MeetingSchema = "meeting"

	// ProjectSchema represents the "project" schema type.
	ProjectSchema = "project"

	// ReviewSchema represents the "review" schema type.
	ReviewSchema = "review"

	// ScratchSchema represents the "scratch" schema type.
	ScratchSchema = "scratch"

	// SystemSchema represents the "system" schema type.
	SystemSchema = "system"
)

// Schemas returns a slice of all valid known schema types.
func Schemas() []string {
	return []string{
		AreaSchema,
		EntitySchema,
		MeetingSchema,
		ProjectSchema,
		ReviewSchema,
		ScratchSchema,
		SystemSchema,
	}
}
