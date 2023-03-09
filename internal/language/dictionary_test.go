package language_test

import (
	"regexp"
	"testing"

	"github.com/revett/atlas/internal/language"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestAllWordlists(t *testing.T) {
	t.Parallel()

	t.Run("elements are unique", func(t *testing.T) {
		t.Parallel()

		all := append([]string{}, language.Animals()...)
		all = append(all, language.Colors()...)
		all = append(all, language.Emotions()...)
		require.ElementsMatch(t, all, lo.Uniq(all))
	})
}

func TestAnimals(t *testing.T) {
	t.Parallel()

	t.Run("uses kebab case", func(t *testing.T) {
		t.Parallel()

		animals := language.Animals()

		for _, animal := range animals {
			t.Run(animal, func(t *testing.T) {
				kebabCasePattern := regexp.MustCompile(`^[a-z]+(-[a-z]+)*$`)
				require.True(t, kebabCasePattern.MatchString(animal))
			})
		}
	})

	t.Run("elements are unique", func(t *testing.T) {
		t.Parallel()

		animals := language.Animals()
		require.ElementsMatch(t, animals, lo.Uniq(animals))
	})
}

func TestColors(t *testing.T) {
	t.Parallel()

	t.Run("uses single words only", func(t *testing.T) {
		t.Parallel()

		colors := language.Colors()

		for _, color := range colors {
			t.Run(color, func(t *testing.T) {
				ok, err := regexp.MatchString("^[a-z]+$", color)
				require.NoError(t, err)
				require.True(t, ok)
			})
		}
	})

	t.Run("elements are unique", func(t *testing.T) {
		t.Parallel()

		colors := language.Colors()
		require.ElementsMatch(t, colors, lo.Uniq(colors))
	})
}

func TestEmotions(t *testing.T) {
	t.Parallel()

	t.Run("uses single words only", func(t *testing.T) {
		t.Parallel()

		emotions := language.Emotions()

		for _, emotion := range emotions {
			t.Run(emotion, func(t *testing.T) {
				ok, err := regexp.MatchString("^[a-z]+$", emotion)
				require.NoError(t, err)
				require.True(t, ok)
			})
		}
	})

	t.Run("elements are unique", func(t *testing.T) {
		t.Parallel()

		emotions := language.Emotions()
		require.ElementsMatch(t, emotions, lo.Uniq(emotions))
	})
}
