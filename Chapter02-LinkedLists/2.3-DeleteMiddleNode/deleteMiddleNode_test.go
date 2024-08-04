package delete_middle_node

import (
	"CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"
	"testing"
)

// Test case for deleting a middle node
func TestDeleteMiddleNode(t *testing.T) {
	head := LinkedList.CreateList([]int{1, 2, 3, 4, 5})
	middleNode := head.Next.Next // Node with value 3

	DeleteMiddleNode(middleNode)

	expected := []int{1, 2, 4, 5}
	expectedLinkedList := LinkedList.CreateList(expected)
	if !LinkedList.ListsEqual(head, expectedLinkedList) {
		t.Errorf("DeleteMiddleNode failed. Expected: %v, Got: %v", expected, head)
	}
}

// Test case for deleting the first node (the head)
func TestDeleteMiddleNode_FirstNode(t *testing.T) {
	head := LinkedList.CreateList([]int{1, 2, 3})
	DeleteMiddleNode(head)

	expected := []int{2, 3}
	expectedLinkedList := LinkedList.CreateList(expected)
	if !LinkedList.ListsEqual(head, expectedLinkedList) {
		t.Errorf("DeleteMiddleNode failed. Expected: %v, Got: %v", expected, head)
	}
}

// Test case where the node to delete is the last node (should do nothing)
func TestDeleteMiddleNode_LastNode(t *testing.T) {
	head := LinkedList.CreateList([]int{1, 2, 3})
	lastNode := head.Next.Next // Node with value 3

	DeleteMiddleNode(lastNode)

	expected := []int{1, 2, 3} // Should remain unchanged
	expectedLinkedList := LinkedList.CreateList(expected)

	if !LinkedList.ListsEqual(head, expectedLinkedList) {
		t.Errorf("DeleteMiddleNode failed. Expected: %v, Got: %v", expected, head)
	}
}

// Test case where the node is nil (should do nothing)
func TestDeleteMiddleNode_NilNode(t *testing.T) {
	head := LinkedList.CreateList([]int{1, 2, 3})
	var nilNode *LinkedList.ListNode

	DeleteMiddleNode(nilNode)

	expected := []int{1, 2, 3} // Should remain unchanged
	expectedLinkedList := LinkedList.CreateList(expected)

	if !LinkedList.ListsEqual(head, expectedLinkedList) {
		t.Errorf("DeleteMiddleNode failed. Expected: %v, Got: %v", expected, head)
	}
}
