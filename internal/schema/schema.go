package schema

// TODO: refactor to be types with .Stringer interface, as well as filename
// generation func.

const (
	ArchiveSchema = "archive"
	AreaSchema    = "area"
	EntitySchema  = "entity"
	MeetingSchema = "meeting"
	ProjectSchema = "project"
	ScratchSchema = "scratch"
	SystemSchema  = "system"
)

// Schemas returns a slice of all valid known schema types.
func Schemas() []string {
	return []string{
		AreaSchema,
		EntitySchema,
		MeetingSchema,
		ProjectSchema,
		ScratchSchema,
		SystemSchema,
	}
}
