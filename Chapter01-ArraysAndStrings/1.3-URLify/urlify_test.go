package urlify

import (
	"testing"
)

func TestURLify(t *testing.T) {
	input := "Mr John Smith    "
	trueLength := 13
	expected := "Mr%20John%20Smith"

	result := URLify(input, trueLength)

	if result != expected {
		t.Errorf("URLify(%q, %d) = %q; expected %q", input, trueLength, result, expected)
	} else {
		t.Logf("URLify(%q, %d) = %q; expected %q", input, trueLength, result, expected)
	}
}
