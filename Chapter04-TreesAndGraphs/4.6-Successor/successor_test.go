package successor

import (
	"fmt"
	"testing"
)

// Helper function to construct a sample binary search tree
func constructBST() *Node {
	root := &Node{Val: 20}
	root.Left = &Node{Val: 10, Parent: root}
	root.Right = &Node{Val: 30, Parent: root}
	root.Left.Left = &Node{Val: 5, Parent: root.Left}
	root.Left.Right = &Node{Val: 15, Parent: root.Left}
	root.Right.Right = &Node{Val: 35, Parent: root.Right}
	root.Left.Left.Left = &Node{Val: 3, Parent: root.Left.Left}
	root.Left.Left.Right = &Node{Val: 7, Parent: root.Left.Left}
	root.Left.Right.Left = &Node{Val: 13, Parent: root.Left.Right}
	root.Left.Right.Left.Left = &Node{Val: 14, Parent: root.Left.Right.Left}
	root.Left.Right.Right = &Node{Val: 17, Parent: root.Left.Right}
	return root
}

// Helper function to print the tree structure in a human-readable format
func printTree(root *Node) {
	if root == nil {
		return
	}
	printTreeHelper(root, "", true)
}

// Helper function to print tree nodes recursively
func printTreeHelper(node *Node, indent string, last bool) {
	if node == nil {
		return
	}

	fmt.Print(indent)
	if last {
		fmt.Print("R---- ")
		indent += "     "
	} else {
		fmt.Print("L---- ")
		indent += "|    "
	}

	fmt.Println(node.Val)

	printTreeHelper(node.Left, indent, false)
	printTreeHelper(node.Right, indent, true)
}

// TestFindSuccessor tests the FindSuccessor method for various nodes
func TestFindSuccessor(t *testing.T) {
	root := constructBST()

	// Print the full tree at the start of the test
	fmt.Println("Binary Search Tree:")
	printTree(root)

	tests := []struct {
		nodeValue int
		expected  int
	}{
		{10, 14}, // In-order successor of 10 is 14
		{15, 17}, // In-order successor of 15 is 17
		{17, 20}, // In-order successor of 17 is 20
		{20, 30}, // In-order successor of 20 is 30
		{30, 35}, // In-order successor of 30 is 35
		{35, -1}, // In-order successor of 35 is nil (no successor)
		{7, 10},  // In-order successor of 7 is 10
		{3, 5},   // In-order successor of 3 is 5
	}

	for _, test := range tests {
		node := findNode(root, test.nodeValue)
		successor := node.FindSuccessor()
		if successor == nil {
			if test.expected != -1 {
				t.Errorf("Node with value %d should have a successor with value %d, but found nil", test.nodeValue, test.expected)
			} else {
				t.Logf("Node with value %d has no successor as expected", test.nodeValue)
			}
		} else if successor.Val != test.expected {
			t.Errorf("Node with value %d: expected successor value %d, but got %d", test.nodeValue, test.expected, successor.Val)
		} else {
			t.Logf("Node with value %d: found expected successor value %d", test.nodeValue, test.expected)
		}
	}
}

// Helper function to find a node by value in the BST
func findNode(root *Node, value int) *Node {
	if root == nil {
		return nil
	}
	if root.Val == value {
		return root
	}
	if value < root.Val {
		return findNode(root.Left, value)
	}
	return findNode(root.Right, value)
}
