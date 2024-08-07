package flip_bit_to_win

import (
	"testing"
)

func TestFlipBitToWin(t *testing.T) {
	tests := []struct {
		num      int64
		expected int
	}{
		{
			num:      1775, // Binary: 11011101111, max sequence by flipping a 0: 8
			expected: 8,
		},
		{
			num:      0b1111, // Binary: 1111, no 0s to flip, max sequence: 4
			expected: 4,
		},
		{
			num:      0b1000, // Binary: 1000, flip one 0 to get 1100, max sequence: 2
			expected: 2,
		},
		{
			num:      0b101010, // Binary: 101010, flip one 0 to get 111010 or 101110, max sequence: 3
			expected: 3,
		},
		{
			num:      0b11101110011, // Binary: 11101110011, flip one 0 to get 11111110011, max sequence: 8
			expected: 7,
		},
		{
			num:      0b0, // Binary: 0, flip one 0 to get 1, max sequence: 1
			expected: 1,
		},
		{
			num:      0b1, // Binary: 1, no 0s to flip, max sequence: 1
			expected: 1,
		},
		{
			num:      0b11111, // Binary: 11111, no 0s to flip, max sequence: 1
			expected: 5,
		},
		{
			num:      0b011100111110111110001111110010, // Binary: 0b011100111110111110001111110010, flip one 0 to get 0b011100111110111110001111110010, max sequence: 11
			expected: 11,
		},
		{
			num:      0b1011001000110110, // Binary: 1011001000110110, flip one 0 to get 1011001000111110, max sequence: 6
			expected: 5,
		},
		{
			num:      0b1000000000000000, // Binary: 1000000000000000, flip one 0 to get 1100000000000000, max sequence: 2
			expected: 2,
		},
	}

	for _, test := range tests {
		result := FlipBitToWin(test.num)
		if result != test.expected {
			t.Errorf("Num %d FlipBitToWin(%b) = %d; expected %d", test.num, test.num, result, test.expected)
		} else {
			t.Logf("Num %d FlipBitToWin(%b) = %d; expected %d [PASS]", test.num, test.num, result, test.expected)
		}
	}
}
