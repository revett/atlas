package validate_test

import (
	"testing"

	"github.com/revett/atlas/internal/validate"
	"github.com/stretchr/testify/require"
)

func Test_filenameValidator_Validate(t *testing.T) { //nolint:funlen
	t.Parallel()

	tests := map[string]struct {
		input string
		errs  []error
	}{
		"SuccessWithoutUsingDashes": {
			input: "area.language.go.interfaces.md",
		},
		"SuccessUsingDashes": {
			input: "area.language.go.error-handling.md",
		},
		"SuccessNumeric": {
			input: "area.2022.36.md",
		},
		"SuccessNumericWithDashes": {
			input: "scratch.2022.05.18.13-25-27.md",
		},
		"SuccessWithNumericMixed": {
			input: "meeting.2022.06.08.1432.interview.technical.md",
		},
		"SuccessArchiveSchema": {
			input: "archive.project.video-app.md",
		},
		"SuccessAreaSchema": {
			input: "area.language.go.generics.md",
		},
		"SuccessEntitySchema": {
			input: "entity.person.business.bill-gates.md",
		},
		"SuccessProjectSchema": {
			input: "project.video-app.md",
		},
		"SuccessScratchSchema": {
			input: "scratch.2022.05.18.132527.md",
		},
		"SuccessSystemSchema": {
			input: "system.checklist.monthly-accounts.md",
		},
		"ErrorInvalidBaseSchema": {
			input: "interview.2022.06.08.1432.md",
			errs: []error{
				validate.ErrUnrecognisedBaseSchemaType,
			},
		},
		"ErrorUppercaseAll": {
			input: "AREA.LANGUAGE.GO.ERRORS.md",
			errs: []error{
				validate.ErrUnrecognisedBaseSchemaType,
				validate.ErrInvalidFilenameFormat,
			},
		},
		"ErrorUppercaseStart": {
			input: "AREA.language.go.errors.md",
			errs: []error{
				validate.ErrUnrecognisedBaseSchemaType,
				validate.ErrInvalidFilenameFormat,
			},
		},
		"ErrorUppercaseEnd": {
			input: "area.language.go.ERRORS.md",
			errs: []error{
				validate.ErrInvalidFilenameFormat,
			},
		},
		"ErrorUppercaseSome": {
			input: "area.lanGuage.go.errors.md",
			errs: []error{
				validate.ErrInvalidFilenameFormat,
			},
		},

		"ErrorTooManyDots": {
			input: "area..language.go.errors.md",
			errs: []error{
				validate.ErrInvalidFilenameFormat,
			},
		},
		"ErrorSpaceInMiddle": {
			input: "area.language.go.error handling.md",
			errs: []error{
				validate.ErrInvalidFilenameFormat,
			},
		},
		"ErrorSpaceAtStart": {
			input: " area.language.go.errors.md",
			errs: []error{
				validate.ErrUnrecognisedBaseSchemaType,
				validate.ErrInvalidFilenameFormat,
			},
		},
		"ErrorMissingFileExtension": {
			input: "area.2022.36",
			errs: []error{
				validate.ErrMissingMarkdownFileExtension,
				validate.ErrInvalidFilenameFormat,
			},
		},
		"ErrorOnlySchemaProvided": {
			input: "area.md",
			errs: []error{
				validate.ErrInvalidFilenameFormat,
				validate.ErrInsufficientNumberOfFilenameParts,
			},
		},
		"ErrorDotAtEnd": {
			input: "area.language.go.errors.",
			errs: []error{
				validate.ErrMissingMarkdownFileExtension,
				validate.ErrInvalidFilenameFormat,
			},
		},
		"ErrorDotAtEndAfterFileExtension": {
			input: "area.language.go.errors.md.",
			errs: []error{
				validate.ErrMissingMarkdownFileExtension,
				validate.ErrInvalidFilenameFormat,
			},
		},
		"ErrorDotAtStart": {
			input: ".area.language.go.errors.md",
			errs: []error{
				validate.ErrUnrecognisedBaseSchemaType,
				validate.ErrInvalidFilenameFormat,
			},
		},
		"ErrorSpaceAtEnd": {
			input: "area.language.go.errors.md ",
			errs: []error{
				validate.ErrMissingMarkdownFileExtension,
				validate.ErrInvalidFilenameFormat,
			},
		},
	}

	for n, testCase := range tests {
		tc := testCase

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			errs := validate.NewFilenameValidator().Validate(tc.input)

			require.Len(t, errs, len(tc.errs))
			for i, err := range errs {
				require.ErrorIs(t, err, tc.errs[i])
			}
		})
	}
}
