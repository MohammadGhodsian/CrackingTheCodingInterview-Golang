package urlify

/*
URLify:
Write a method to replace all spaces in a string with '%20:
You may assume that the string has sufficient space at the end to hold the additional characters,
and that you are given the "true" length of the string. (Note: If implementing in Java,
please use a character array so that you can perform this operation in place.)
EXAMPLE
Input: "Mr John Smith " J 13
Output: "Mr%20J ohn%20Smith"
*/

// URLify replaces all spaces in a string with '%20:
func URLify(str string, trueLength int) string {
	runes := []rune(str)
	// Count the number of spaces within the true length
	spaceCount := 0
	for i := 0; i < trueLength; i++ {
		if runes[i] == ' ' {
			spaceCount++
		}
	}

	// Calculate the new length
	newLength := trueLength + spaceCount*2
	if newLength > len(runes) {
		return "" // Not enough space
	}

	// Modify the string in place from the end
	index := newLength - 1
	for i := trueLength - 1; i >= 0; i-- {
		if runes[i] == ' ' {
			runes[index] = '0'
			runes[index-1] = '2'
			runes[index-2] = '%'
			index -= 3
		} else {
			runes[index] = runes[i]
			index--
		}
	}
	return string(runes)
}
