package palindrome_permutation

import (
	"unicode"
)

/*
Palindrome Permutation:
Given a string, write a function to check if it is a permutation of a palindrome.
A palindrome is a word or phrase that is the same forwards and backwards.
A permutation is a rearrangement of letters.
The palindrome does not need to be limited to just dictionary words.
EXAMPLE
Input: Tact Coa
Output: True (permutations: "taco cat". "atco cta". etc.)
*/

func IsPalindromePermutation(given string) bool {
	charCount := make(map[rune]int)
	for _, char := range given {
		// If we need to check type of characters, we can use comething like: if unicode.IsLetter(char) || unicode.IsDigit(char)
		if char != ' ' { // Ignore spaces
			char = unicode.ToLower(char)
			charCount[char]++
		}
	}
	oddCount := 0
	for _, count := range charCount {
		if count%2 == 1 {
			oddCount++
		}
		if oddCount > 1 {
			return false
		}
	}
	return true
}
