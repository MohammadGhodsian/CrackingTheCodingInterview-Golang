package sum_lists

import (
	"CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"
	"testing"
)

func TestSumListsReverseOrder(t *testing.T) {
	// Test case 1
	l1 := LinkedList.CreateList([]int{7, 1, 6})
	l2 := LinkedList.CreateList([]int{5, 9, 2})
	expected := LinkedList.CreateList([]int{2, 1, 9})

	result := SumListsReverseOrder(l1, l2)
	if !LinkedList.ListsEqual(result, expected) {
		t.Errorf("SumListsReverseOrder failed. Expected: %v, Got: %v", expected, result)
	}
}

func TestSumListsForwardOrder(t *testing.T) {
	// Test case 1
	l1 := LinkedList.CreateList([]int{6, 1, 7})
	l2 := LinkedList.CreateList([]int{2, 9, 5})
	expected := LinkedList.CreateList([]int{9, 1, 2})

	result := SumListsForwardOrder(l1, l2)
	if !LinkedList.ListsEqual(result, expected) {
		t.Errorf("SumListsForwardOrder failed. Expected: %v, Got: %v", expected, result)
	}
}
