package list_of_depths

/*
4.3 List of Depths
Given a binary tree, design an algorithm which creates a linked list of all the nodes at each depth
(e.g., if you have a tree with depth 0, you'll have 0 linked lists).
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ListOfDepths returns a map where the key is the depth of the tree and
// the value is a slice of TreeNode pointers that are at that depth.
func (t *TreeNode) ListOfDepths() map[int][]*TreeNode {
	// Create a map to store nodes at each depth level
	mp := make(map[int][]*TreeNode)

	// Start the depth-first search with the root node at depth 0
	listOfDepthsHelper(t, 0, mp)

	return mp
}

// listOfDepthsHelper is a recursive helper function that populates the
// map with nodes at each depth.
func listOfDepthsHelper(tn *TreeNode, depth int, mp map[int][]*TreeNode) {
	// Base case: If the current node is nil, return
	if tn == nil {
		return
	}

	// Append the current node to the list corresponding to its depth in the map
	mp[depth] = append(mp[depth], tn)

	// Recursively process the left and right children, increasing the depth by 1
	listOfDepthsHelper(tn.Left, depth+1, mp)
	listOfDepthsHelper(tn.Right, depth+1, mp)
}
