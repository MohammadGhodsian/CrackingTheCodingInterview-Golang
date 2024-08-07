package route_between_nodes

/*
4.1 Route Between Nodes
Given a directed graph, design an algorithm to find out whether there is a route between two nodes.
*/

type Graph struct {
	vertices map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(from, to int) {
	g.vertices[from] = append(g.vertices[from], to)
}

// DFS performs a Depth-First Search to determine if there is a path from 'from' to 'to'.
// It returns true if such a path exists and false otherwise.
func (g Graph) DFS(from, to int) bool {
	// Map to keep track of visited nodes to prevent revisiting and potential infinite loops.
	visited := make(map[int]bool)
	// Start the DFS helper function to search for a path from 'from' to 'to'.
	return dfsHelper(g, from, to, visited)
}

// dfsHelper is a recursive function that performs the Depth-First Search.
// It explores all paths from the current node to the target node 'to'.
// visited: keeps track of nodes that have already been visited in the current search.
func dfsHelper(g Graph, current, to int, visited map[int]bool) bool {
	// If the current node is the destination node, return true.
	if current == to {
		return true
	}
	// Mark the current node as visited.
	visited[current] = true
	// Explore each neighbor of the current node.
	for _, neighbor := range g.vertices[current] {
		// If the neighbor has not been visited, recursively search from the neighbor.
		if !visited[neighbor] {
			if dfsHelper(g, neighbor, to, visited) {
				return true
			}
		}
	}
	// If no path was found, return false.
	return false
}
