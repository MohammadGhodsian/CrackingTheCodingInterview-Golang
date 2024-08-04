package oneAway

import (
	"math"
)

/*
One Away:
There are three types of edits that can be performed on strings:
insert a character, remove a character, or replace a character.
Given two strings, write a function to check if they are one edit (or zero edits) away.
EXAMPLE
pale, ple -> true
pales, pale -> true
pale, bale -> true
pale, bake -> false
*/

func OneEditAway(first, second string) bool {
	ln1, ln2 := len(first), len(second)
	// If the length difference is greater than 1, more than one edit is required.
	if math.Abs(float64(ln1-ln2)) > 1 {
		return false
	}

	var shorter, longer string
	// Identify the shorter and longer string
	if ln1 < ln2 {
		shorter, longer = first, second
	} else {
		shorter, longer = second, first
	}
	ln1 = len(shorter)
	ln2 = len(longer)

	i, j := 0, 0
	foundDifference := false
	// Iterate through both strings
	for i < ln1 && j < ln2 {
		if shorter[i] == longer[j] {
			i++ // Move shorter string pointer if characters match
		} else {
			if foundDifference {
				return false // More than one difference found
			}
			foundDifference = true // Mark that a difference was found
			if ln1 == ln2 {
				i++ // Move shorter string pointer if lengths are the same
			}
		}
		j++ // Always move longer string pointer
	}
	return true
}
