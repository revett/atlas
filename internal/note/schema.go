package note

const (
	SystemSchema    = "system"
	ProjectSchema   = "project"
	EntitySchema    = "entity"
	InterviewSchema = "interview"
	AreaSchema      = "area"
	ScratchSchema   = "scratch"
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
