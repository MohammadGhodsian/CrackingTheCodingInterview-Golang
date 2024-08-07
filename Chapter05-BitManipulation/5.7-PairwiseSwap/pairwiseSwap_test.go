package pairwise_swap

import (
	"testing"
)

func TestPairwiseSwapOddEvenBits(t *testing.T) {
	tests := []struct {
		input    uint32
		expected uint32
	}{
		{0b1010, 0b0101},         // 10 -> 5
		{0b1111, 0b1111},         // 15 -> 15
		{0b1001, 0b0110},         // 9  -> 6
		{0b110011, 0b110011},     // 51 -> 51
		{0b110, 0b1001},          // 6  -> 9
		{0xFFFFFFFF, 0xFFFFFFFF}, // All bits set -> All bits set
		{0xAAAAAAAA, 0x55555555}, // Alternating bits starting with 1 -> Alternating bits starting with 0
		{0x55555555, 0xAAAAAAAA}, // Alternating bits starting with 0 -> Alternating bits starting with 1
	}

	for _, test := range tests {
		result := PairwiseSwapOddEvenBits(test.input)
		if result != test.expected {
			t.Errorf("PairwiseSwapOddEvenBits(%032b) = %032b; expected %032b", test.input, result, test.expected)
		} else {
			t.Logf("PairwiseSwapOddEvenBits(%032b) = %032b; expected %032b [PASS]", test.input, result, test.expected)
		}
	}
}
