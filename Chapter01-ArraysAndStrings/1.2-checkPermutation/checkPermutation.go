package check_permutation

import (
	"sort"
)

/*
Check Permutation:
Given two strings, write a method to decide if one is a permutation of the other.
*/

func IsPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	// Convert the strings to runes
	r1 := []rune(s1)
	r2 := []rune(s2)
	// Sort the runes
	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})
	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})
	// Compare the sorted strings
	return string(r1) == string(r2)
}
