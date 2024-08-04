package is_unique

import (
	"testing"
)

func TestHasUniqueCharacters(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"abcdefg", true},
		{"hello", false},
		{"1234567890", true},
		{"GoLang", true},
		{"Unique", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := HasUniqueCharacters(tc.input)
			if actual != tc.expected {
				t.Errorf("hasUniqueCharacters(%q) = %v; expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestHasUniqueCharactersWithoutDataStructures(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"abcdefg", true},
		{"hello", false},
		{"1234567890", true},
		{"GoLang", true},
		{"Unique", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := HasUniqueCharactersWithoutDataStructures(tc.input)
			if actual != tc.expected {
				t.Errorf("hasUniqueCharactersWithoutDataStructures(%q) = %v; expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}
