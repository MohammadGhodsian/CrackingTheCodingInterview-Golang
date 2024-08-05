package intersection

import (
	"CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"
	"testing"
)

// Helper function to find the node at a specific index in a linked list
func getNodeAt(head *LinkedList.ListNode, index int) *LinkedList.ListNode {
	current := head
	for i := 0; i < index && current != nil; i++ {
		current = current.Next
	}
	return current
}

func TestFindIntersection_ListsIntersect(t *testing.T) {
	list1 := LinkedList.CreateList([]int{1, 2, 3, 4, 5})
	list2 := LinkedList.CreateList([]int{9, 8})
	intersectingNode := getNodeAt(list1, 2) // Node with value 3
	list2.Next.Next = intersectingNode      // Intersect at node with value 3

	result := FindIntersection(list1, list2)
	if result != intersectingNode {
		t.Errorf("TestFindIntersection_ListsIntersect failed: expected %v, got %v", intersectingNode, result)
	}
}

func TestFindIntersection_ListsDoNotIntersect(t *testing.T) {
	list1 := LinkedList.CreateList([]int{1, 2, 3})
	list2 := LinkedList.CreateList([]int{9, 8, 7})

	result := FindIntersection(list1, list2)
	if result != nil {
		t.Errorf("TestFindIntersection_ListsDoNotIntersect failed: expected nil, got %v", result)
	}
}

func TestFindIntersection_OneListEmpty(t *testing.T) {
	var list1 *LinkedList.ListNode
	list2 := LinkedList.CreateList([]int{1, 2, 3})

	result := FindIntersection(list1, list2)
	if result != nil {
		t.Errorf("TestFindIntersection_OneListEmpty failed: expected nil, got %v", result)
	}
}

func TestFindIntersection_BothListsEmpty(t *testing.T) {
	var list1, list2 *LinkedList.ListNode

	result := FindIntersection(list1, list2)
	if result != nil {
		t.Errorf("TestFindIntersection_BothListsEmpty failed: expected nil, got %v", result)
	}
}

func TestFindIntersection_ListsIntersectAtHead(t *testing.T) {
	list1 := LinkedList.CreateList([]int{1, 2, 3})
	list2 := list1 // Both lists point to the same head node

	result := FindIntersection(list1, list2)
	if result != list1 {
		t.Errorf("TestFindIntersection_ListsIntersectAtHead failed: expected %v, got %v", list1, result)
	}
}
