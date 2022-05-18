package note

const (
	SystemType noteType = iota
	ProjectType
	EntityType
	InterviewType
	AreaType
	ScratchType
)

type noteType int

// String implements the Stringer.String() interface.
func (n noteType) String() string {
	switch n {
	case SystemType:
		return "system"
	case ProjectType:
		return "project"
	case EntityType:
		return "entity"
	case InterviewType:
		return "interview"
	case AreaType:
		return "area"
	case ScratchType:
		return "scratch"
	default:
		return "unknown"
	}
}

func Types() []string {
	return []string{
		SystemType.String(),
		ProjectType.String(),
		EntityType.String(),
		InterviewType.String(),
		AreaType.String(),
		ScratchType.String(),
	}
}
