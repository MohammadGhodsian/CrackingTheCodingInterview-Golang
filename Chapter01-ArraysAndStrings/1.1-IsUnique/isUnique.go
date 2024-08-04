package isUnique

import (
	"sort"
)

/*
Implement an algorithm to determine if a string has all unique characters.
What if you cannot use additional data structures?
*/

// HasUniqueCharacters checks if a string has all unique characters.
func HasUniqueCharacters(s string) bool {
	// Create a map to keep track of characters we've seen.
	seen := make(map[rune]bool)

	// Iterate over each character in the string.
	for _, char := range s {
		// If we've seen the character before, return false.
		if seen[char] {
			return false
		}
		// Mark the character as seen.
		seen[char] = true
	}

	// If we get through the string without duplicates, return true.
	return true
}

// HasUniqueCharactersWithoutDataStructures checks if a string has all unique characters without using additional data structures.
func HasUniqueCharactersWithoutDataStructures(s string) bool {
	// Convert string to a slice of runes for sorting.
	runes := []rune(s)

	// Sort the slice of runes.
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Check for adjacent duplicates.
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			return false
		}
	}

	return true
}
