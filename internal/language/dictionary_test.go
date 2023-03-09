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
		unique := lo.Uniq(animals)
		require.Equal(t, len(animals), len(unique))
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
		unique := lo.Uniq(colors)
		require.Equal(t, len(colors), len(unique))
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
		unique := lo.Uniq(emotions)
		require.Equal(t, len(emotions), len(unique))
	})
}
