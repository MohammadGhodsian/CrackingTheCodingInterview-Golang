package list_of_depths

/*
4.3 List of Depths
Given a binary tree, design an algorithm which creates a linked list of all the nodes at each depth
(e.g., if you have a tree with depth 0, you'll have 0 linked lists).
*/

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ListNode represents a node in the linked list that stores TreeNode pointers.
type ListNode struct {
	*TreeNode
	Next *ListNode
}

// ListOfDepths returns a map where the key is the depth of the tree and
// the value is the head of a linked list containing TreeNode pointers at that depth.
func (t *TreeNode) ListOfDepths() map[int]*ListNode {
	// Create a map to store linked lists of nodes at each depth level
	mp := make(map[int]*ListNode)

	// Start the depth-first search with the root node at depth 0
	listOfDepthsHelper(t, 0, mp)

	return mp
}

// listOfDepthsHelper is a recursive helper function that populates the
// map with linked lists of nodes at each depth.
func listOfDepthsHelper(tn *TreeNode, depth int, mp map[int]*ListNode) {
	// Base case: If the current node is nil, return
	if tn == nil {
		return
	}

	// Create a new list node
	newNode := ListNode{
		TreeNode: tn,
	}
	// Find the saved node at current depth in map
	n, ok := mp[depth]
	if n == nil || !ok {
		// If map has no value, add the created list node to map
		mp[depth] = &newNode
	} else {
		// Find last item in linked list
		for n.Next != nil {
			n = n.Next
		}
		// Move the create list node to end of linked list
		n.Next = &newNode
	}

	// Recursively process the left and right children, increasing the depth by 1
	listOfDepthsHelper(tn.Left, depth+1, mp)
	listOfDepthsHelper(tn.Right, depth+1, mp)
}
