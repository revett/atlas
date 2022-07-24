package input_test

import (
	"testing"

	"github.com/revett/sepias/internal/input"
	"github.com/stretchr/testify/require"
)

func TestValidateTitleFormat(t *testing.T) { // nolint:funlen
	t.Parallel()

	tests := map[string]struct {
		input   string
		errFunc func(*testing.T, error)
	}{
		"Simple": {
			input: "foo.bar.baz",
			errFunc: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"SimpleKebab": {
			input: "foo.bar-baz.qux",
			errFunc: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"SimpleNumeric": {
			input: "foo.123.baz",
			errFunc: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"SimpleNumericStart": {
			input: "123.bar.baz",
			errFunc: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"SimpleNumericEnd": {
			input: "foo.bar.123",
			errFunc: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"ComplexNumeric": {
			input: "foo.123-456.baz",
			errFunc: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"UppercaseAll": {
			input: "FOO.BAR.BAZ",
			errFunc: func(t *testing.T, err error) {
				require.ErrorIs(t, err, input.ErrInvalidTitleFormat)
			},
		},
		"UppercaseStart": {
			input: "FOO.bar.baz",
			errFunc: func(t *testing.T, err error) {
				require.ErrorIs(t, err, input.ErrInvalidTitleFormat)
			},
		},
		"UppercaseEnd": {
			input: "foo.bar.BAZ",
			errFunc: func(t *testing.T, err error) {
				require.ErrorIs(t, err, input.ErrInvalidTitleFormat)
			},
		},
		"DotAtEnd": {
			input: "foo.bar.baz.",
			errFunc: func(t *testing.T, err error) {
				require.ErrorIs(t, err, input.ErrInvalidTitleFormat)
			},
		},
		"DotAtStart": {
			input: ".foo.bar.baz",
			errFunc: func(t *testing.T, err error) {
				require.ErrorIs(t, err, input.ErrInvalidTitleFormat)
			},
		},
		"DoubleDot": {
			input: "foo..bar.baz",
			errFunc: func(t *testing.T, err error) {
				require.ErrorIs(t, err, input.ErrInvalidTitleFormat)
			},
		},
		"Space": {
			input: "foo. bar.baz",
			errFunc: func(t *testing.T, err error) {
				require.ErrorIs(t, err, input.ErrInvalidTitleFormat)
			},
		},
		"SpaceStart": {
			input: " foo.bar.baz",
			errFunc: func(t *testing.T, err error) {
				require.ErrorIs(t, err, input.ErrInvalidTitleFormat)
			},
		},
	}

	for n, testCase := range tests {
		tc := testCase // nolint:varnamelen

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			err := input.ValidateTitleFormat(tc.input)
			tc.errFunc(t, err)
		})
	}
}

func TestValidateTitleBaseSchemaType(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input   string
		errFunc func(*testing.T, error)
	}{
		"Unknown": {
			input: "foo.bar.baz",
			errFunc: func(t *testing.T, err error) {
				require.ErrorIs(t, err, input.ErrUnrecognisedBaseSchemaTypeErr)
			},
		},
		"Empty": {
			input: "",
			errFunc: func(t *testing.T, err error) {
				require.ErrorIs(t, err, input.ErrInsufficientNumberOfTitlePartsErr)
			},
		},
		"SingleValidPart": {
			input: "area",
			errFunc: func(t *testing.T, err error) {
				require.ErrorIs(t, err, input.ErrInsufficientNumberOfTitlePartsErr)
			},
		},
		"Success": {
			input: "area.language.go.errors",
			errFunc: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"ArchiveType": {
			input: "archive.area.language.go.errors",
			errFunc: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
	}

	for n, testCase := range tests {
		tc := testCase // nolint:varnamelen

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			err := input.ValidateTitleBaseSchemaType(tc.input)
			tc.errFunc(t, err)
		})
	}
}
