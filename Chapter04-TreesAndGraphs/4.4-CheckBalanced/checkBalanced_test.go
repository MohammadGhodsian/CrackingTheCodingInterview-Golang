package check_balanced

import (
	"testing"
)

// Test cases
func TestCheckBalanced(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "Empty Tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Single Node",
			root:     &TreeNode{Val: 1},
			expected: true,
		},
		{
			name: "Balanced Tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 6},
				},
			},
			expected: true,
		},
		{
			name: "Unbalanced Tree (Left-heavy)",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 8},
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: false,
		},
		{
			name: "Unbalanced Tree (Right-heavy)",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val:   5,
						Right: &TreeNode{Val: 6},
					},
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.root.CheckBalanced()
			if result != tt.expected {
				t.Errorf("CheckBalanced() = %v, want %v", result, tt.expected)
			}
			t.Logf("Test %s: got %v, expected %v", tt.name, result, tt.expected)
		})
	}
}
