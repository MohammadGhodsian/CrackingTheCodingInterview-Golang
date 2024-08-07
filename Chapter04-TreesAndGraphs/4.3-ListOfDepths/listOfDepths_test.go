package list_of_depths

import (
	"reflect"
	"testing"
)

// Helper function to convert linked list to values
func linkedListToValues(head *ListNode) []int {
	values := []int{}
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Val)
	}
	return values
}

// Test cases
func TestListOfDepths(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected map[int][]int
	}{
		{
			name:     "Empty Tree",
			root:     nil,
			expected: map[int][]int{},
		},
		{
			name: "Single Node",
			root: &TreeNode{Val: 1},
			expected: map[int][]int{
				0: {1},
			},
		},
		{
			name: "Simple Tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: map[int][]int{
				0: {1},
				1: {2, 3},
			},
		},
		{
			name: "Tree with Multiple Depths",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: map[int][]int{
				0: {1},
				1: {2, 3},
				2: {4, 5, 6},
			},
		},
		{
			name: "Deep Tree with Single Child",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
					},
				},
			},
			expected: map[int][]int{
				0: {1},
				1: {2},
				2: {3},
				3: {4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s\n", tt.name)
			got := tt.root.ListOfDepths()
			gotMap := make(map[int][]int)
			for depth, list := range got {
				gotMap[depth] = linkedListToValues(list)
			}

			// Log the results
			t.Logf("Expected: %+v\n", tt.expected)
			t.Logf("Got: %+v\n", gotMap)

			if !reflect.DeepEqual(gotMap, tt.expected) {
				t.Errorf("ListOfDepths() = %v, want %v", gotMap, tt.expected)
			}
		})
	}
}
