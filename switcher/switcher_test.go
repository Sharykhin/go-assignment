package switcher

import (
	"io"
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

	tt := []struct {
		name          string
		inA   bool
		inB  bool
		inC bool
		inD float64
		inE int
		inF int
		expectedOut float64
 		expectedError error
	}{
		{
			name: "A && B && !C => H = M. H = M => K = D + (D * E / 10)",
			inA: true,
			inB: true,
			inC: false,
			inD: 10.32,
			inE: 10,
			inF: 10,
			expectedOut: 20.64,
			expectedError: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out, err := l.Calculate(tc.inA, tc.inB, tc.inC, tc.inD, tc.inE, tc.inF, Base)
			assert.Equal(tc.expectedOut, out)
			assert.Equal(tc.expectedError, err)
		})
	}
}