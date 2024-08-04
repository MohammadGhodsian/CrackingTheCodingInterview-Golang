package check_permutation

import (
	"testing"
)

func TestIsPermutation(t *testing.T) {
	testCases := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"abc", "bca", true},
		{"hello", "billion", false},
		{"rat", "tar", true},
		{"night", "thing", true},
		{"abcd", "dcba", true},
	}

	for _, tc := range testCases {
		t.Run(tc.s1+"_"+tc.s2, func(t *testing.T) {
			actual := IsPermutation(tc.s1, tc.s2)
			if actual != tc.expected {
				t.Errorf("IsPermutation(%q, %q) = %v; expected %v", tc.s1, tc.s2, actual, tc.expected)
			}
		})
	}
}
