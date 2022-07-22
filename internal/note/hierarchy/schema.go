package hierarchy

const (
	// ArchiveSchema represents the "archive" schema type.
	ArchiveSchema = "archive"

	// AreaSchema represents the "area" schema type.
	AreaSchema = "area"

	// EntitySchema represents the "entity" schema type.
	EntitySchema = "entity"

	// InterviewSchema represents the "interview" schema type.
	InterviewSchema = "interview"

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
		InterviewSchema,
		ProjectSchema,
		ReviewSchema,
		ScratchSchema,
		SystemSchema,
	}
}
