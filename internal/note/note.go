package note

import "fmt"

type Note struct {
	Schema string
}

func NewNote(s string) (Note, error) {
	for _, e := range Schemas() {
		if e == s {
			return Note{
				Schema: s,
			}, nil
		}
	}

	return Note{}, fmt.Errorf("unknown schema: %s", s)
}
