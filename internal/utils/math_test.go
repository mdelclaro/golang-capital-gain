package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"golang-capital-gain/internal/utils"
)

func TestApplyPrecision(t *testing.T) {
	tests := []struct {
		name      string
		value     float64
		precision int
		expected  float64
	}{
		{
			name:      "Round to 2 decimal places",
			value:     123.456,
			precision: 2,
			expected:  123.46,
		},
		{
			name:      "Round to 0 decimal places",
			value:     123.456,
			precision: 0,
			expected:  123,
		},
		{
			name:      "Round to 3 decimal places",
			value:     123.4567,
			precision: 3,
			expected:  123.457,
		},
		{
			name:      "Round negative number to 2 decimal places",
			value:     -123.456,
			precision: 2,
			expected:  -123.46,
		},
		{
			name:      "Round small number to 4 decimal places",
			value:     0.123456,
			precision: 4,
			expected:  0.1235,
		},
		{
			name:      "Round large number to 1 decimal place",
			value:     123456.789,
			precision: 1,
			expected:  123456.8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.ApplyPrecision(tt.value, tt.precision)
			assert.Equal(t, tt.expected, result)
		})
	}
}
