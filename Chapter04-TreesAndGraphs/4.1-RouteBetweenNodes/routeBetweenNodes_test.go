package route_between_nodes

import (
	"testing"
)

func TestGraph_DFS(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(1, 6)
	graph.AddEdge(6, 7)
	graph.AddEdge(7, 8)

	tests := []struct {
		from     int
		to       int
		expected bool
	}{
		{1, 5, true},
		{1, 8, true},
		{2, 5, true},
		{5, 1, false},
		{8, 1, false},
		{3, 1, false},
		{1, 7, true},
		{6, 5, false},
		{6, 8, true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := graph.DFS(tt.from, tt.to)
			if result != tt.expected {
				t.Errorf("DFS(%d, %d) = %v; expected %v", tt.from, tt.to, result, tt.expected)
			} else {
				t.Logf("SUCCESS: DFS(%d, %d) = %v; expected %v", tt.from, tt.to, result, tt.expected)
			}
		})
	}
}
