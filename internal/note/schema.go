package note

const (
	// SystemSchema is a const for representing the "system" schema type.
	SystemSchema = "system"

	// ProjectSchema is a const for representing the "project" schema type.
	ProjectSchema = "project"

	// EntitySchema is a const for representing the "entity" schema type.
	EntitySchema = "entity"

	// InterviewSchema is a const for representing the "interview" schema type.
	InterviewSchema = "interview"

	// AreaSchema is a const for representing the "area" schema type.
	AreaSchema = "area"

	// ScratchSchema is a const for representing the "scratch" schema type.
	ScratchSchema = "scratch"
)

// Schemas returns a slice of all valid known schema types.
func Schemas() []string {
	return []string{
		SystemSchema,
		ProjectSchema,
		EntitySchema,
		InterviewSchema,
		AreaSchema,
		ScratchSchema,
	}
}
