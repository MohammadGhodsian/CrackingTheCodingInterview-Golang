package insertion

/*
5.1 Insertion
You are given two 32-bit numbers, N and M, and two bit positions, i and j.
Write a method to insert M into N such that M starts at bit j and ends at bit i.
You can assume that the bits j through i have enough space to fit all of M. That is, if M=10011,
 you can assume that there are at least 5 bits between j and i.
You would not, for example, have j=3 and i=2,
  because M could not fully fit between bit 3 and bit 2.

EXAMPLE

Input: N = 10000000000, M = 10011, i = 2, j = 6

Output: N = 10001001100
*/

// InsertMIntoN inserts M into N such that M starts at bit j and ends at bit i.
func InsertMIntoN(N, M, i, j int) int {
	// Step 1: Create a mask to clear bits i through j in N
	// Example for i=2, j=6: left = 11100000000, right = 00000000011, mask = 11100000011
	allOnes := ^0              // All bits 1
	left := allOnes << (j + 1) // 1s before position j, then 0s
	right := (1 << i) - 1      // 1s after position i
	mask := left | right       // 1s before position j, and after position i

	// Step 2: Clear the bits j through i in N
	N_cleared := N & mask

	// Step 3: Shift M so that it lines up with bits j through i
	M_shifted := M << i

	// Step 4: Combine N_cleared with M_shifted
	return N_cleared | M_shifted
}
