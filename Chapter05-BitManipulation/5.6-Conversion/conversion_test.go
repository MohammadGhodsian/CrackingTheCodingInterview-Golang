package conversion

import (
	"testing"
)

func TestConversion(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{
			a:        29, // 11101
			b:        15, // 01111
			expected: 2,
		},
		{
			a:        0, // 0000
			b:        0, // 0000
			expected: 0,
		},
		{
			a:        1, // 0001
			b:        0, // 0000
			expected: 1,
		},
		{
			a:        15, // 01111
			b:        0,  // 00000
			expected: 4,
		},
		{
			a:        255, // 11111111
			b:        0,   // 00000000
			expected: 8,
		},
		{
			a:        0,   // 0000
			b:        255, // 11111111
			expected: 8,
		},
		{
			a:        1023, // 1111111111
			b:        511,  // 0111111111
			expected: 1,
		},
	}

	for _, test := range tests {
		result := Conversion(test.a, test.b)
		if result != test.expected {
			t.Errorf("Conversion(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		} else {
			t.Logf("Conversion(%d, %d) = %d; expected %d [PASS]", test.a, test.b, result, test.expected)
		}
	}
}
