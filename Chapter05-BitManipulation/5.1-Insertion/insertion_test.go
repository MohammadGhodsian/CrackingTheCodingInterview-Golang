package insertion

import (
	"testing"
)

func TestInsertMIntoN(t *testing.T) {
	tests := []struct {
		N, M     int
		i, j     int
		expected int
	}{
		{
			N:        0b10000000000,
			M:        0b10011,
			i:        2,
			j:        6,
			expected: 0b10001001100,
		},
		{
			N:        0b10000000000,
			M:        0b1,
			i:        2,
			j:        2,
			expected: 0b10000000100,
		},
		{
			N:        0b0,
			M:        0b111,
			i:        0,
			j:        2,
			expected: 0b111,
		},
	}

	for _, test := range tests {
		result := InsertMIntoN(test.N, test.M, test.i, test.j)
		if result != test.expected {
			t.Errorf("InsertMIntoN(%b, %b, %d, %d) = %b; expected %b",
				test.N, test.M, test.i, test.j, result, test.expected)
		} else {
			t.Logf("InsertMIntoN(%b, %b, %d, %d) = %b; expected %b [PASS]",
				test.N, test.M, test.i, test.j, result, test.expected)
		}
	}
}
