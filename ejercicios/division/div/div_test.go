package div

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name           string
	A              float64
	B              float64
	ExpectedResult float64
	ExpectedError  bool
}

func TestDivision(t *testing.T) {
	for _, testCase := range []TestCase{
		{
			Name:           "8/4",
			A:              8,
			B:              4,
			ExpectedResult: 2,
			ExpectedError:  false,
		},
		{
			Name:           "9/2",
			A:              9,
			B:              2,
			ExpectedResult: 4.5,
			ExpectedError:  false,
		},
		{
			Name:          "9/0",
			A:             9,
			B:             0,
			ExpectedError: true,
		},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			result, error := Division(testCase.A, testCase.B)
			if testCase.ExpectedError {
				assert.NotNil(t, error)
				return
			}
			assert.EqualValues(t, testCase.ExpectedResult, result)
		})
	}
}
