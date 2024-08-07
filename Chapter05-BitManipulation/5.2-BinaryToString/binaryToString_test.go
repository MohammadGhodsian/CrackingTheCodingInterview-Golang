package binary_to_string

import (
	"testing"
)

func TestFloatToBinary(t *testing.T) {
	tests := []struct {
		num      float32
		expected string
	}{
		{
			num:      0.1,
			expected: "0.000110011001100110011001101",
		},
		{
			num:      0.5,
			expected: "0.1",
		},
		{
			num:      0.25,
			expected: "0.01",
		},
		{
			num:      0.625,
			expected: "0.101",
		},
		{
			num:      0.0,
			expected: "0.0",
		},
		{
			num:      1.0,
			expected: "ERROR", // Edge case: 1 is not between 0 and 1.
		},
		{
			num:      -0.5,
			expected: "ERROR", // Edge case: Negative number is invalid.
		},
		{
			num:      0.75,
			expected: "0.11",
		},
		{
			num:      0.999999,
			expected: "0.111111111111111111101111",
		},
	}

	for _, test := range tests {
		result := FloatToBinary(test.num)
		if result != test.expected {
			t.Errorf("FloatToBinary(%f) = %s; expected %s", test.num, result, test.expected)
		} else {
			t.Logf("FloatToBinary(%f) = %s; expected %s [PASS]", test.num, result, test.expected)
		}
	}
}
