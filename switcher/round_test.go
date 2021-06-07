package switcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		name          string
		in float64
		unit float64
		expectedOut float64
	}{
		{
			name: "Round up",
			in: 10.356,
			unit:0.01,
			expectedOut: 10.36,
		},
		{
			name: "Round down",
			in: 10.353,
			unit:0.01,
			expectedOut: 10.35,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := round(tc.in, tc.unit)
			assert.Equal(tc.expectedOut, out)
		})
	}
}