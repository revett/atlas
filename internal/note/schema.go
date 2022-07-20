package note

const (
	// SystemSchema represents the "system" schema type.
	SystemSchema = "system"

	// ProjectSchema represents the "project" schema type.
	ProjectSchema = "project"

	// EntitySchema represents the "entity" schema type.
	EntitySchema = "entity"

	// InterviewSchema represents the "interview" schema type.
	InterviewSchema = "interview"

	// AreaSchema represents the "area" schema type.
	AreaSchema = "area"

	// ReviewSchema represents the "review" schema type.
	ReviewSchema = "review"

	// ScratchSchema represents the "scratch" schema type.
	ScratchSchema = "scratch"
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
