package oneAway

import (
	"testing"
)

func TestOneEditAway(t *testing.T) {
	testCases := [][3]string{
		{"pale", "ple", "true"},
		{"pales", "pale", "true"},
		{"pale", "bale", "true"},
		{"pale", "bake", "false"},
	}

	for _, pair := range testCases {
		t.Run(pair[0]+"_"+pair[1], func(t *testing.T) {
			actual := OneEditAway(pair[0], pair[1])
			expected := pair[2] == "true"
			if actual != expected {
				t.Errorf("OneEditAway(%q, %q) = %v; expected %v", pair[0], pair[1], actual, expected)
			} else {
				t.Logf("OneEditAway(%q, %q) = %v; expected %v", pair[0], pair[1], actual, expected)
			}
		})
	}
}
