package switcher

import (
	"errors"
	"io"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Sharykhin/go-assignment/logger"
)

func init() {
	logger.Log.SetOutput(io.Discard)
}

func TestLogical_Calculate(t *testing.T) {
	assert := assert.New(t)
	l := NewLogical()
	compareFloats := func(a, b float64) bool {
		tolerance := 0.001
		if diff := math.Abs(a - b); diff < tolerance {
			return true
		} else {
			return false
		}
	}

	tt := []struct {
		name          string
		inA   bool
		inB  bool
		inC bool
		inD float64
		inE int
		inF int
		inMode Mode
		expectedOut float64
 		expectedError error
	}{
		{
			name: "When A && B && !C Then H = M And K = D + (D * E / 10)",
			inA: true,
			inB: true,
			inC: false,
			inD: 10.32,
			inE: 10,
			inF: 10,
			inMode:Base,
			expectedOut: 20.64,
			expectedError: nil,
		},
		{
			name: "When A && B && C Then H = P And K = D + (D * (E - F) / 25.5)",
			inA: true,
			inB: true,
			inC: true,
			inD: 10.32,
			inE: 10,
			inF: 8,
			inMode:Base,
			expectedOut: 11.13,
			expectedError: nil,
		},
		{
			name: "When A && B && C Then H = P And K = D + (D * (E - F) / 25.5)",
			inA: true,
			inB: true,
			inC: true,
			inD: 10.32,
			inE: 100,
			inF: 8,
			inMode:Base,
			expectedOut: 47.55,
			expectedError: nil,
		},
		{
			name: "When !A && B && C Then H = T And K = K = D - (D * F / 30)",
			inA: false,
			inB: true,
			inC: true,
			inD: 10.32,
			inE: 10,
			inF: 8,
			inMode:Base,
			expectedOut: 7.57,
			expectedError: nil,
		},
		{
			name: "When !A && !B && C Then an error is return",
			inA: false,
			inB: false,
			inC: true,
			inD: 10.32,
			inE: 10,
			inF: 8,
			inMode:Base,
			expectedOut: 0,
			expectedError: ErrUnexpectedInput,
		},
		{
			name: "When mode is custom1 And A && B && C Then H = P And K = 2 * D + (D * E / 100)",
			inA: true,
			inB: true,
			inC: true,
			inD: 10.32,
			inE: 10,
			inF: 8,
			inMode:CustomOne,
			expectedOut: 21.67,
			expectedError: nil,
		},
		{
			name: "When mode is custom2 And A && !B && C Then H = M And K = F + D + (D * E / 100)",
			inA: true,
			inB: false,
			inC: true,
			inD: 10.32,
			inE: 10,
			inF: 8,
			inMode:CustomTwo,
			expectedOut: 19.35,
			expectedError: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out, err := l.Calculate(tc.inA, tc.inB, tc.inC, tc.inD, tc.inE, tc.inF, tc.inMode)
			assert.True(compareFloats(tc.expectedOut, out))
			if err != nil {
				assert.True(errors.Is(err, tc.expectedError))
			}
		})
	}
}
