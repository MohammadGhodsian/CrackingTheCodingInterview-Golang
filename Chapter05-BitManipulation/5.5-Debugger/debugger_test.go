package debugger

import (
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		// Positive numbers that are powers of two
		{1, true},   // 2^0
		{2, true},   // 2^1
		{4, true},   // 2^2
		{8, true},   // 2^3
		{16, true},  // 2^4
		{32, true},  // 2^5
		{64, true},  // 2^6
		{128, true}, // 2^7

		// Positive numbers that are not powers of two
		{3, false}, // Not a power of two
		{5, false}, // Not a power of two
		{6, false}, // Not a power of two
		{7, false}, // Not a power of two
		{9, false}, // Not a power of two

		// Edge cases
		{0, true},   // Result of Zero & any number is zero
		{-1, false}, // Negative number
		{-2, false}, // Negative number
		{-4, false}, // Negative number
		{-8, false}, // Negative number
	}

	for _, test := range tests {
		result := IsZeroOrPositivePowerOfTwo(test.n)
		if result != test.expected {
			t.Errorf("IsPowerOfTwo(%d) = %v; expected %v", test.n, result, test.expected)
		} else {
			t.Logf("IsPowerOfTwo(%d) = %v; expected %v [PASS]", test.n, result, test.expected)
		}
	}
}
