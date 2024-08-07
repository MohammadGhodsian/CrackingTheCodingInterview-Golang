package validate_bst

import (
	"fmt"
	"testing"
)

// Helper function to print the tree vertically
func printTree(node *BinaryTreeNode, space int, height int) string {
	if node == nil {
		return ""
	}

	// Increase distance between levels
	space += height

	// Process right child first
	right := printTree(node.Right, space, height)

	// Current node after spaces
	result := fmt.Sprintf("%s%d\n", generateSpaces(space-height), node.Val)

	// Process left child
	left := printTree(node.Left, space, height)

	return right + result + left
}

// Helper function to generate spaces for tree printing
func generateSpaces(n int) string {
	return fmt.Sprintf("%*s", n, "")
}

func treeToString(root *BinaryTreeNode) string {
	return printTree(root, 0, 10)
}

func TestIsBST(t *testing.T) {
	// Test case 1: A valid BST
	root1 := &BinaryTreeNode{Val: 10}
	root1.Left = &BinaryTreeNode{Val: 5}
	root1.Right = &BinaryTreeNode{Val: 15}
	root1.Left.Left = &BinaryTreeNode{Val: 2}
	root1.Left.Right = &BinaryTreeNode{Val: 7}
	root1.Right.Left = &BinaryTreeNode{Val: 12}
	root1.Right.Right = &BinaryTreeNode{Val: 20}
	root1.Right.Right.Right = &BinaryTreeNode{Val: 22}

	if !root1.IsBST() {
		t.Errorf("Test case 1 failed: Expected true, got false. Tree structure:\n%s", treeToString(root1))
	} else {
		t.Logf("Test case 1 passed: Tree is a valid BST. Tree structure:\n%s", treeToString(root1))
	}

	// Test case 2: An invalid BST (left child greater than root)
	root2 := &BinaryTreeNode{Val: 10}
	root2.Left = &BinaryTreeNode{Val: 15}
	root2.Right = &BinaryTreeNode{Val: 20}

	if root2.IsBST() {
		t.Errorf("Test case 2 failed: Expected false, got true. Tree structure:\n%s", treeToString(root2))
	} else {
		t.Logf("Test case 2 passed: Tree is not a valid BST. Tree structure:\n%s", treeToString(root2))
	}

	// Test case 3: An invalid BST (right child less than root)
	root3 := &BinaryTreeNode{Val: 10}
	root3.Left = &BinaryTreeNode{Val: 5}
	root3.Right = &BinaryTreeNode{Val: 7}

	if root3.IsBST() {
		t.Errorf("Test case 3 failed: Expected false, got true. Tree structure:\n%s", treeToString(root3))
	} else {
		t.Logf("Test case 3 passed: Tree is not a valid BST. Tree structure:\n%s", treeToString(root3))
	}

	// Test case 4: A single node tree (valid BST)
	root4 := &BinaryTreeNode{Val: 10}

	if !root4.IsBST() {
		t.Errorf("Test case 4 failed: Expected true, got false. Tree structure:\n%s", treeToString(root4))
	} else {
		t.Logf("Test case 4 passed: Single node tree is a valid BST. Tree structure:\n%s", treeToString(root4))
	}

	// Test case 5: An empty tree (valid BST)
	var root5 *BinaryTreeNode

	if !root5.IsBST() {
		t.Errorf("Test case 5 failed: Expected true, got false. Tree structure:\n%s", treeToString(root5))
	} else {
		t.Logf("Test case 5 passed: Empty tree is a valid BST. Tree structure:\n%s", treeToString(root5))
	}
}
