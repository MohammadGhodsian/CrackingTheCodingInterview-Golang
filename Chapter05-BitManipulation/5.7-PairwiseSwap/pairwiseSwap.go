package pairwise_swap

/*
5.7 Pairwise Swap
Write a program to swap odd and even bits in an integer with as few instructions as possible
 (e.g., bit 0 and bit 1 are swapped, bit 2 and bit 3 are swapped, and so on).
*/

func PairwiseSwapOddEvenBits(n uint32) uint32 {
	// Masks for odd and even bit positions
	evenMask := uint32(0xAAAAAAAA) // 10101010... for 32-bit
	oddMask := uint32(0x55555555)  // 01010101... for 32-bit

	// Extract and shift even bits to odd positions
	evenBits := (n & evenMask) >> 1

	// Extract and shift odd bits to even positions
	oddBits := (n & oddMask) << 1

	// Combine the shifted bits
	return evenBits | oddBits
}
