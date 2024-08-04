package string_compression

import (
	"testing"
)

func TestCompress(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"abcdefgh", "abcdefgh"},
		{"aabcccccaaa", "a2b1c5a3"},
		{"aaaabccdddeeee", "a4b1c2d3e4"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := Compress(tc.input)
			if actual != tc.expected {
				t.Errorf("compress(%q) = %q; expected %q", tc.input, actual, tc.expected)
			} else {
				t.Logf("compress(%q) = %q; expected %q", tc.input, actual, tc.expected)
			}
		})
	}
}
