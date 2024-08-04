package rotateMatrix

// Rotate Matrix:
// Given an image represented by an NxN matrix,
// where each pixel in the image is 4 bytes,
// write a method to rotate the image by 90 degrees.
// can you do this in place?

type Matrix [][]int

func RotateMatrix(matrix Matrix) Matrix {
	ln := len(matrix)
	for concentric := 0; concentric < ln/2; concentric++ {
		frstConcentricIndex := concentric
		lastConcentricIndex := ln - 1 - concentric
		for concentricIndex := frstConcentricIndex; concentricIndex < lastConcentricIndex; concentricIndex++ {
			offset := concentricIndex - frstConcentricIndex
			matrix[concentricIndex][lastConcentricIndex],
				matrix[frstConcentricIndex][concentricIndex],
				matrix[lastConcentricIndex-offset][frstConcentricIndex],
				matrix[lastConcentricIndex][lastConcentricIndex-offset] =
				matrix[frstConcentricIndex][concentricIndex],
				matrix[lastConcentricIndex-offset][frstConcentricIndex],
				matrix[lastConcentricIndex][lastConcentricIndex-offset],
				matrix[concentricIndex][lastConcentricIndex]
			/*
				// save the top element in current iteration
				temp := matrix[frstConcentricIndex][concentricIndex]
				// keep the current offset
				offset := concentricIndex - frstConcentricIndex
				// move left element to top
				matrix[frstConcentricIndex][concentricIndex] = matrix[lastConcentricIndex-offset][frstConcentricIndex]
				// move bottom element to left
				matrix[lastConcentricIndex-offset][frstConcentricIndex] = matrix[lastConcentricIndex][lastConcentricIndex-offset]
				// move right element to bottom
				matrix[lastConcentricIndex][lastConcentricIndex-offset] = matrix[concentricIndex][lastConcentricIndex]
				// assign saved temp (top element) to right
				matrix[concentricIndex][lastConcentricIndex] = temp
			*/
		}
	}
	return matrix
}
