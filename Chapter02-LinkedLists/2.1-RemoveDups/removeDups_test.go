package remove_dups

import (
	linkedList "CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"
	"testing"
)

// Test case where there are no duplicates
func TestRemoveDuplicatesNoBuffer_NoDuplicates(t *testing.T) {
	head := linkedList.CreateList([]int{1, 2, 3, 4, 5})
	expected := linkedList.CreateList([]int{1, 2, 3, 4, 5})

	result := RemoveDuplicatesNoBuffer(head)
	if !linkedList.ListsEqual(result, expected) {
		t.Errorf("RemoveDuplicatesNoBuffer failed. Expected: %v, Got: %v", expected, result)
	}
}

// Test case where all elements are duplicates
func TestRemoveDuplicatesNoBuffer_AllDuplicates(t *testing.T) {
	head := linkedList.CreateList([]int{2, 2, 2, 2, 2})
	expected := linkedList.CreateList([]int{2})

	result := RemoveDuplicatesNoBuffer(head)
	if !linkedList.ListsEqual(result, expected) {
		t.Errorf("RemoveDuplicatesNoBuffer failed. Expected: %v, Got: %v", expected, result)
	}
}

// Test case where there are scattered duplicates
func TestRemoveDuplicatesNoBuffer_ScatteredDuplicates(t *testing.T) {
	head := linkedList.CreateList([]int{1, 2, 2, 3, 4, 4, 5})
	expected := linkedList.CreateList([]int{1, 2, 3, 4, 5})

	result := RemoveDuplicatesNoBuffer(head)
	if !linkedList.ListsEqual(result, expected) {
		t.Errorf("RemoveDuplicatesNoBuffer failed. Expected: %v, Got: %v", expected, result)
	}
}

// Test case where the list is empty
func TestRemoveDuplicatesNoBuffer_EmptyList(t *testing.T) {
	head := linkedList.CreateList([]int{})
	expected := linkedList.CreateList([]int{})

	result := RemoveDuplicatesNoBuffer(head)
	if !linkedList.ListsEqual(result, expected) {
		t.Errorf("RemoveDuplicatesNoBuffer failed. Expected: %v, Got: %v", expected, result)
	}
}
