package successor

/*
4.6 Successor
Write an algorithm to find the "next" node (i.e., in-order successor) of a given node in a binary search tree.
You may assume that each node has a link to its parent.
*/

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

// FindSuccessor finds the in-order successor of the current node in a binary search tree.
func (n *Node) FindSuccessor() *Node {
	// Case 1: Node has a right child
	if n.Right != nil {
		// Go to the right child
		s := n.Right
		// Find the leftmost node in the right subtree
		for s.Left != nil {
			s = s.Left
		}
		return s
	}
	// Case 2: Node does not have a right child
	// Traverse up using the parent pointers to find the successor
	current := n       // Start with the current node
	parent := n.Parent // Initialize parent to the current node's parent
	// Traverse up the tree
	for parent != nil && current != parent.Left {
		// Move up to the parent node
		current = parent
		parent = parent.Parent
	}
	// Return the parent node as the successor if found
	// If the loop exits because parent is nil, it means the current node is the rightmost node
	// and hence does not have an in-order successor, so return nil
	return parent
}
