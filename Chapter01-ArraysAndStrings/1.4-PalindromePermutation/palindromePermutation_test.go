package palindromePermutation

import (
	"testing"
)

func TestIsPalindromePermutation(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"Tact Coa", true},
		{"racecar", true},
		{"apple", false},
		{"carerac", true},
		{"A Santa at NASA", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := IsPalindromePermutation(tc.input)
			if actual != tc.expected {
				t.Errorf("IsPalindromePermutation(%q) = %v; expected %v", tc.input, actual, tc.expected)
			} else {
				t.Logf("IsPalindromePermutation(%q) = %v; expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}
