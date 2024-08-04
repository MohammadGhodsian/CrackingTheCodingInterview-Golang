package string_compression

import "fmt"

/*
String Compression:
Implement a method to perform basic string compression using the counts of repeated characters.
For example, the string aabcccccaaa would become a2b1c5a3.
If the "compressed" string would not become smaller than the original string, your method should return the original string.
You can assume the string has only uppercase and lowercase letters (a - z).
*/

func Compress(s string) string {
	if len(s) == 0 {
		return s
	}

	lastChar := rune(s[0]) // Initialize the last character to the first character of the string
	count := 0             // Initialize the count of the current character sequence
	result := ""           // Initialize the result string

	for _, char := range s {
		if char == lastChar {
			count++ // Increment the count if the current character matches the last character
		} else {
			// Append the last character and its count to the result string
			result += fmt.Sprintf("%c%d", lastChar, count)
			lastChar = char // Update the last character to the current character
			count = 1       // Reset the count to 1 for the new character
		}
	}
	// Append the final character and its count to the result string
	result += fmt.Sprintf("%c%d", lastChar, count)

	// Return the original string if the compressed version is not shorter
	if len(result) < len(s) {
		return result
	}
	return s
}
