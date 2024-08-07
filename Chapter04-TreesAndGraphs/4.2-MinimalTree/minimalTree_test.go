package minimal_tree

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestBuildBST(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect *TreeNode
	}{
		{
			name:   "Empty array",
			input:  []int{},
			expect: nil,
		},
		{
			name:   "Single element array",
			input:  []int{1},
			expect: &TreeNode{Val: 1},
		},
		{
			name:  "Two elements array",
			input: []int{1, 2},
			expect: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
		},
		{
			name:  "Three elements array",
			input: []int{1, 2, 3},
			expect: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			name:  "Multiple elements array",
			input: []int{1, 2, 3, 4, 5, 6, 7},
			expect: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   6,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			name:  "Unsorted array",
			input: []int{7, 1, 5, 2, 6, 3, 4},
			expect: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   6,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 7},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildBST(tt.input)
			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("BuildBST(%v) =\n%v\nwant\n%v", tt.input, treeToString(result), treeToString(tt.expect))
			} else {
				t.Logf("Success. Given list: %+v\nResult tree:\n%v", tt.input, treeToString(result))
			}
		})
	}
}

func treeToString(node *TreeNode) string {
	if node == nil {
		return "nil"
	}
	var buffer bytes.Buffer
	printTree(&buffer, node, 0)
	return buffer.String()
}

func printTree(buffer *bytes.Buffer, node *TreeNode, level int) {
	if node == nil {
		return
	}
	printTree(buffer, node.Right, level+1)
	buffer.WriteString(fmt.Sprintf("%s%d\n", indent(level), node.Val))
	printTree(buffer, node.Left, level+1)
}

func indent(level int) string {
	return string(bytes.Repeat([]byte("  "), level))
}
