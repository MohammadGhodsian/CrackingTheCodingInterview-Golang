package check_balanced

import "math"

/*
4.4 Check Balanced
Implement a function to check if a binary tree is balanced.
For the purposes of this question, a balanced tree is defined to be a tree such that the heights of the two subtrees of any node never differ by more than one.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CheckBalanced determines if the binary tree rooted at this node is balanced.
// A balanced tree is defined as a tree where the heights of the two subtrees of any node
// never differ by more than one.
func (t *TreeNode) CheckBalanced() bool {
	_, balanced := checkBalancedHelper(t)
	return balanced
}

// checkBalancedHelper is a helper function that computes the height of the subtree rooted at `t`
// and determines whether the subtree is balanced.
// It returns two values:
// 1. The height of the subtree rooted at `t`.
// 2. A boolean indicating whether the subtree is balanced.
func checkBalancedHelper(t *TreeNode) (height int, balanced bool) {
	if t == nil {
		// An empty subtree is balanced and has height 0.
		return 0, true
	}

	// Recursively check the left subtree and the right subtree.
	leftHeight, leftBalanced := checkBalancedHelper(t.Left)
	rightHeight, rightBalanced := checkBalancedHelper(t.Right)

	// Compute the height of the current subtree.
	if leftHeight > rightHeight {
		height = leftHeight
	} else {
		height = rightHeight
	}
	height++

	// Determine if the current subtree is balanced.
	balanced =
		leftBalanced && // The left subtree must be balanced.
			rightBalanced && // The right subtree must be balanced.
			math.Abs(float64(leftHeight-rightHeight)) <= 1 // The height difference between subtrees must be at most 1.
	return
}
