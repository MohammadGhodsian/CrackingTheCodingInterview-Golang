package rotateMatrix

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func printMatrix(matrix Matrix) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func TestRotateMatrix(t *testing.T) {
	type sample struct {
		Input    Matrix
		Expected Matrix
	}

	testCases := []sample{
		{
			Input: Matrix{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			Expected: Matrix{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			Input: Matrix{
				{0, 1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10, 11},
				{12, 13, 14, 15, 16, 17},
				{18, 19, 20, 21, 22, 23},
				{24, 25, 26, 27, 28, 29},
				{30, 31, 32, 33, 34, 35},
			},
			Expected: Matrix{
				{30, 24, 18, 12, 6, 0},
				{31, 25, 19, 13, 7, 1},
				{32, 26, 20, 14, 8, 2},
				{33, 27, 21, 15, 9, 3},
				{34, 28, 22, 16, 10, 4},
				{35, 29, 23, 17, 11, 5},
			},
		},
	}

	for i, test := range testCases {
		t.Run("RotateMatrix "+strconv.Itoa(i), func(t *testing.T) {
			t.Log("Input matrix:")
			printMatrix(test.Input)
			t.Log("Expected matrix:")
			printMatrix(test.Expected)
			rotated := RotateMatrix(test.Input)
			t.Log("Rotated matrix:")
			printMatrix(rotated)
			if !reflect.DeepEqual(rotated, test.Expected) {
				t.Error("FAIL\n")
			} else {
				t.Log("OK\n")
			}
		})
	}

}
