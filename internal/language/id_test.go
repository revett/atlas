package language_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/revett/atlas/internal/language"
	"github.com/stretchr/testify/require"
)

func TestHumanReadableID(t *testing.T) {
	t.Parallel()

	t.Run("has minimum three parts", func(t *testing.T) {
		t.Parallel()

		id := language.HumanReadableID()
		parts := strings.Split(id, "-")
		require.True(t, len(parts) >= 3)
	})

	t.Run("uses kebab case", func(t *testing.T) {
		t.Parallel()

		id := language.HumanReadableID()
		kebabCasePattern := regexp.MustCompile(`^[a-z]+(-[a-z]+)*$`)
		require.True(t, kebabCasePattern.MatchString(id))
	})

	t.Run("generates different ids each time", func(t *testing.T) {
		t.Parallel()

		idOne := language.HumanReadableID()
		idTwo := language.HumanReadableID()
		require.False(t, idOne == idTwo)
	})
}
