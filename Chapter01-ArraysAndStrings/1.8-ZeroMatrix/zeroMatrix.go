package zeroMatrix

/*
Zero Matrix:
Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to O.
*/

type Matrix [][]int

func ZeroMatrix(matrix Matrix) Matrix {
	// Create sets to track rows and columns to be zeroed
	zeroRows, zeroCols := []int{}, []int{}
	// First pass: identify rows and columns to be zeroed
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				zeroRows = append(zeroRows, i)
				zeroCols = append(zeroCols, j)
			}
		}
	}
	// Second pass: set rows to zero
	for _, row := range zeroRows {
		for j := 0; j < len(matrix[row]); j++ {
			matrix[row][j] = 0
		}
	}
	// Second pass: set columns to zero
	for _, col := range zeroCols {
		for i := 0; i < len(matrix); i++ {
			matrix[i][col] = 0
		}
	}
	return matrix
}
