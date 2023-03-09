package language

import (
	"fmt"
	"math/rand"
	"strings"
)

// HumanReadableID generates a random human readable ID in the format of "emotion-color-animal", it
// will always use kebab-case.
func HumanReadableID() string {
	emotions := Emotions()
	colors := Colors()
	animals := Animals()

	id := fmt.Sprintf(
		"%s-%s-%s",
		emotions[rand.Intn(len(emotions))],
		colors[rand.Intn(len(colors))],
		animals[rand.Intn(len(animals))],
	)

	return strings.ToLower(id)
}
