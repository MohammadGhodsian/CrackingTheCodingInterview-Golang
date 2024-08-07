package debugger

/*
 5.5 Debugger
Explain what the following code does: `(n & (n - 1)) == 0`.
*/

// Check if 'n' is zero or a power of two.
// The result of (zero & any number) always is zero.
// A number is a power of two if it has exactly one bit set to 1 in its binary representation.
// The expression 'n & (n - 1)' will be 0 if and only if 'n' is a power of two.
// This works because a power of two in binary is of the form '1000...0',
// and subtracting 1 flips all the bits after the single 1 bit to 1, resulting in a binary number
// where 'n & (n - 1)' clears the only 1 bit in 'n', leaving 0.
//
// Examples:
//
// 1. n = 0 (binary(32-bit): 00000000 00000000 00000000 00000000)
//    n - 1 = -1 (binary(32-bit): 11111111 11111111 11111111 11111111)
//    n & (n−1) = 00000000000000000000000000000000 & 111111111111111111111111111111 = 00000000000000000000000000000000
//
// 2. n = 1 (binary: 0001)
//    n - 1 = 0 (binary: 0000)
//    n & (n-1) = 0001 & 0000 = 0000 (which is 0)
//    Hence, 1 is a power of two.
//
// 3. n = 8 (binary: 1000)
//    n - 1 = 7 (binary: 0111)
//    n & (n-1) = 1000 & 0111 = 0000 (which is 0)
//    Hence, 8 is a power of two.
//
// 4. n = 10 (binary: 1010)
//    n - 1 = 9 (binary: 1001)
//    n & (n-1) = 1010 & 1001 = 1000 (which is not 0)
//    Hence, 10 is not a power of two.
//
// 5. n = -8 (binary: 11111111 11111111 11111111 11111000)
//    n - 1 = -9 (binary: 11111111 11111111 11111111 11110111)
//    n & (n-1) = 11111000 & 11110111 = 11110000 (which is not 0)
//    Hence, -8 is not a power of two.

func IsZeroOrPositivePowerOfTwo(n int) bool {
	return (n & (n - 1)) == 0
}
