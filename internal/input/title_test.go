package input_test

import (
	"testing"

	"github.com/revett/sepias/internal/input"
	"github.com/stretchr/testify/require"
)

func TestValidateTitle(t *testing.T) { // nolint:funlen
	t.Parallel()

	tests := map[string]struct {
		input   string
		errFunc require.ErrorAssertionFunc
	}{
		"Simple": {
			input:   "foo.bar.baz",
			errFunc: require.NoError,
		},
		"SimpleKebab": {
			input:   "foo.bar-baz.qux",
			errFunc: require.NoError,
		},
		"SimpleNumeric": {
			input:   "foo.123.baz",
			errFunc: require.NoError,
		},
		"SimpleNumericStart": {
			input:   "123.bar.baz",
			errFunc: require.NoError,
		},
		"SimpleNumericEnd": {
			input:   "foo.bar.123",
			errFunc: require.NoError,
		},
		"ComplexNumeric": {
			input:   "foo.123-456.baz",
			errFunc: require.NoError,
		},
		"UppercaseAll": {
			input:   "FOO.BAR.BAZ",
			errFunc: require.Error,
		},
		"UppercaseStart": {
			input:   "FOO.bar.baz",
			errFunc: require.Error,
		},
		"UppercaseEnd": {
			input:   "foo.bar.BAZ",
			errFunc: require.Error,
		},
		"DotAtEnd": {
			input:   "foo.bar.baz.",
			errFunc: require.Error,
		},
		"DotAtStart": {
			input:   ".foo.bar.baz",
			errFunc: require.Error,
		},
		"DoubleDot": {
			input:   "foo..bar.baz",
			errFunc: require.Error,
		},
		"Space": {
			input:   "foo. bar.baz",
			errFunc: require.Error,
		},
		"SpaceStart": {
			input:   " foo.bar.baz",
			errFunc: require.Error,
		},
	}

	for n, testCase := range tests {
		tc := testCase // nolint:varnamelen

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			err := input.ValidateTitle(tc.input)
			tc.errFunc(t, err)
		})
	}
}
