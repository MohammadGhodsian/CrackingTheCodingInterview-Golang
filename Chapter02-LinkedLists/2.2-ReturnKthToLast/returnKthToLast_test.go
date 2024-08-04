package return_kth_to_last

import (
	"CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"
	"testing"
)

// Test case for k within the list range
func TestReturnKthToLast_WithinRange(t *testing.T) {
	head := LinkedList.CreateList([]int{1, 2, 3, 4, 5})
	k := 2
	expected := 3

	node, _ := ReturnKthToLast(head, k)
	if node == nil || expected != node.Value {
		t.Errorf("ReturnKthToLast failed. Expected: %v, Got: %v", expected, node)
	}
}

// Test case for k equal to the length of the list - 1 (should return the head)
func TestReturnKthToLast_EqualToLength(t *testing.T) {
	head := LinkedList.CreateList([]int{1, 2, 3, 4, 5})
	k := 4
	expected := 1

	node, _ := ReturnKthToLast(head, k)
	if node == nil || expected != node.Value {
		t.Errorf("ReturnKthToLast failed. Expected: %v, Got: %v", expected, node)
	}
}

// Test case for k equal to 0 (should return the last node)
func TestReturnKthToLast_LastElement(t *testing.T) {
	head := LinkedList.CreateList([]int{1, 2, 3, 4, 5})
	k := 0
	expected := 5

	node, _ := ReturnKthToLast(head, k)
	if node == nil || expected != node.Value {
		t.Errorf("ReturnKthToLast failed. Expected: %v, Got: %v", expected, node)
	}
}

// Test case for k greater than the length of the list (should handle this case gracefully)
func TestReturnKthToLast_GreaterThanLength(t *testing.T) {
	head := LinkedList.CreateList([]int{1, 2, 3})
	k := 5

	node, _ := ReturnKthToLast(head, k)
	if node != nil {
		t.Errorf("ReturnKthToLast failed. Expected: nil, Got: %v", node)
	}
}

// Test case for empty list
func TestReturnKthToLast_EmptyList(t *testing.T) {
	head := LinkedList.CreateList([]int{})
	k := 1

	node, _ := ReturnKthToLast(head, k)
	if node != nil {
		t.Errorf("ReturnKthToLast failed. Expected: nil, Got: %v", node)
	}
}
