package validate_bst

import "math"

/*
4.5 Validate BST
Implement a function to check if a binary tree is a binary search tree.
*/

// BinaryTreeNode represents a node in a binary tree.
type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// IsBST checks if the binary tree rooted at the node is a binary search tree.
func (b *BinaryTreeNode) IsBST() bool {
	// Use helper function with initial min and max values.
	return bstCheckerHelper(b, math.MinInt, math.MaxInt)
}

// bstCheckerHelper is a recursive helper function to validate the BST properties.
func bstCheckerHelper(b *BinaryTreeNode, min, max int) bool {
	// An empty node is considered a valid BST.
	if b == nil {
		return true
	}
	// The current node's value must be strictly between min and max.
	if b.Val <= min || b.Val >= max {
		return false
	}
	// Recursively check the left and right subtrees with updated boundaries.
	return bstCheckerHelper(b.Left, min, b.Val) && bstCheckerHelper(b.Right, b.Val, max)
}
