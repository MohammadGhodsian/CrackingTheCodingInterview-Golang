package stringRotation

import "testing"

func TestCheckStringRotation(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		// Test cases where s2 is a valid rotation of s1
		{"waterbottle", "erbottlewat", true},
		{"hello", "llohe", true},
		{"abcdef", "cdefab", true},
		{"rotation", "tationro", true},

		// Test cases where s2 is not a rotation of s1
		{"hello", "world", false},
		{"abcdef", "abcfed", false},
		{"rotation", "rotations", false},

		// Edge cases
		{"", "", true},                       // Both strings are empty
		{"a", "a", true},                     // Single character strings
		{"a", "b", false},                    // Single character strings, different
		{"abcdefgh", "abcdefghijklm", false}, // Different length strings
	}

	for _, tt := range tests {
		t.Run(tt.s1+"_"+tt.s2, func(t *testing.T) {
			result := CheckStringRotation(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("CheckStringRotation(%q, %q) = %v; expected %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}
