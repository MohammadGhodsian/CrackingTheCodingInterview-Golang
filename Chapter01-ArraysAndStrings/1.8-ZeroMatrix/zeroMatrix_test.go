package zeroMatrix

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	type sample struct {
		Input    Matrix
		Expected Matrix
	}

	tests := []sample{
		{
			Input: Matrix{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			Expected: Matrix{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
		},
		{
			Input: Matrix{
				{1, 1, 1, 1, 1},
				{1, 1, 0, 1, 0},
				{1, 1, 1, 1, 1},
			},
			Expected: Matrix{
				{1, 1, 0, 1, 0},
				{0, 0, 0, 0, 0},
				{1, 1, 0, 1, 0},
			},
		},
		{
			Input: Matrix{
				{1, 0, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 1},
			},
			Expected: Matrix{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 1, 1, 1, 1},
				{0, 0, 1, 1, 1, 1},
				{0, 0, 1, 1, 1, 1},
				{0, 0, 0, 0, 0, 0},
			},
		},
		{
			Input: Matrix{
				{0, 0, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 0, 1, 1},
				{1, 1, 1, 1, 1, 1},
			},
			Expected: Matrix{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 1, 1},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 1, 1},
			},
		},
		{
			Input: Matrix{
				{1, 1, 0},
				{1, 0, 1},
				{0, 1, 1},
			},
			Expected: Matrix{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}

	for _, test := range tests {
		// Create a copy of the input matrix to avoid modifying the original
		input := make(Matrix, len(test.Input))
		for i := range test.Input {
			input[i] = make([]int, len(test.Input[i]))
			copy(input[i], test.Input[i])
		}

		// Apply the function to the input matrix
		ZeroMatrix(input)

		if !reflect.DeepEqual(input, test.Expected) {
			t.Errorf("\nZeroMatrix: %v\n  expected: %v\n    result: %v", test.Input, test.Expected, input)
		} else {
			t.Logf("\nZeroMatrix: %v\n  expected: %v\n    result: %v", test.Input, test.Expected, input)
		}
	}
}
