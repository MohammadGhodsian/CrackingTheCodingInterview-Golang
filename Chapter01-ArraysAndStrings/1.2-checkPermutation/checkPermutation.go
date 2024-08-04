package checkPermutation

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
	r1 := []rune(s1)
	r2 := []rune(s2)

	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})
	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})
	return string(r1) == string(r2)
}
