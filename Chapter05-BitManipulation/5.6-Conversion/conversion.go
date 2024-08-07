package conversion

/*
5.6 Conversion
Write a function to determine the number of bits you would need to flip to convert integer A to integer B.

EXAMPLE
Input: 29 (or: 11101), 15 (or: 01111)
Output: 2
*/

func Conversion(a, b int) (count int) {
	xOr := a ^ b
	// using golang builtin:
	// import "math/bits"
	// return bits.OnesCount(uint(xOr))
	for xOr != 0 { // count 1's
		xOr = xOr & (xOr - 1)
		count++
	}
	return
}
