package minimal_tree

import "sort"

/*
4.2 Minimal Tree
Given a sorted (increasing order) array with unique integer elements,
 write an algo- rithm to create a binary search tree with minimal height.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildBST creates a binary search tree with minimal height from a sorted array
func BuildBST(arr []int) *TreeNode {
	// Create a copy of the input array to avoid modifying the original array
	copySort := make([]int, len(arr))
	copy(copySort, arr)
	// Sort the array to ensure it is in increasing order
	sort.Ints(copySort)
	// Call the recursive helper function to build the BST
	return buildBST(copySort, 0, len(arr)-1)
}

// buildBST is a recursive helper function to build a BST from a sorted array
func buildBST(arr []int, startIndex, endIndex int) *TreeNode {
	// Base case: if the start index is greater than the end index, return nil
	if startIndex < 0 || endIndex >= len(arr) || endIndex < startIndex {
		return nil
	}
	// Find the middle index of the current subarray
	middleIndex := (startIndex + endIndex) / 2
	// Create a new TreeNode with the value at the middle index
	return &TreeNode{
		Val:   arr[middleIndex],
		Left:  buildBST(arr, startIndex, middleIndex-1), // Recursively build the left subtree
		Right: buildBST(arr, middleIndex+1, endIndex),   // Recursively build the right subtree
	}
}
