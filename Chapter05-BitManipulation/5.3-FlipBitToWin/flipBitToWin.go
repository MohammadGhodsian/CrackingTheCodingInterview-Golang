package flip_bit_to_win

import (
	"strconv"
	"strings"
)

/*
5.3 Flip Bit to Win
You have an integer and you can flip exactly one bit from a 0 to a 1.
Write code to find the length of the longest sequence of 1s you could create.
EXAMPLE
Input: 1775 (or: 11011101111)
Output: 8
*/

func FlipBitToWin(num int64) int {
	// Edge case: If num is 0, flipping the only 0 will give a sequence of length 1
	if num == 0 {
		return 1
	}

	// Convert the integer to its binary representation as a string
	binary := strconv.FormatInt(num, 2)

	// Split the binary string into segments of consecutive '1's, separated by '0's
	spliteds := strings.Split(binary, "0")

	// If there are no '0's in the binary string, return the length of the binary string
	if len(spliteds) == 1 {
		return len(binary)
	}

	maxLen := 0
	// Iterate over segments of consecutive '1's
	for i := 1; i < len(spliteds); i++ {
		// Get the length of the current segment and the previous segment
		s := spliteds[i]
		sPrev := spliteds[i-1]
		lnS := len(s)
		lnSp := len(sPrev)

		// If both segments are empty, skip to the next iteration
		if lnS == 0 && lnSp == 0 {
			continue
		}

		// Calculate the potential maximum length by combining the lengths of the two segments and adding 1 for the flipped '0'
		if lnS+lnSp+1 > maxLen {
			maxLen = lnS + lnSp + 1
		}
	}

	// Return the maximum length found
	return maxLen
}
