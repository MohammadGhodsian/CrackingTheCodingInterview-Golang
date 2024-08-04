package partition

import (
	"CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"
	"testing"
)

// Helper function to convert linked list to slice for easy comparison
func listToSlice(head *LinkedList.ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Value)
		head = head.Next
	}
	return result
}

// Test case for partitioning around a value x
func TestPartition(t *testing.T) {
	head := LinkedList.CreateList([]int{3, 5, 8, 5, 10, 2, 1})
	x := 5
	expected := []int{3, 2, 1, 5, 8, 5, 10}

	result := Partition(head, x)
	actual := listToSlice(result)
	if !equalSlices(expected, actual) {
		t.Errorf("Partition failed. Expected: %v, Got: %v", expected, actual)
	}
}

// Test case for partitioning when x is smaller than all nodes
func TestPartition_XSmaller(t *testing.T) {
	head := LinkedList.CreateList([]int{6, 7, 8, 9})
	x := 5
	expected := []int{6, 7, 8, 9}

	result := Partition(head, x)
	actual := listToSlice(result)
	if !equalSlices(expected, actual) {
		t.Errorf("Partition failed. Expected: %v, Got: %v", expected, actual)
	}
}

// Test case for partitioning when x is larger than all nodes
func TestPartition_XLarger(t *testing.T) {
	head := LinkedList.CreateList([]int{1, 2, 3, 4})
	x := 5
	expected := []int{1, 2, 3, 4}

	result := Partition(head, x)
	actual := listToSlice(result)
	if !equalSlices(expected, actual) {
		t.Errorf("Partition failed. Expected: %v, Got: %v", expected, actual)
	}
}

// Test case for an empty list
func TestPartition_EmptyList(t *testing.T) {
	var head *LinkedList.ListNode
	x := 5
	expected := []int{}

	result := Partition(head, x)
	actual := listToSlice(result)
	if !equalSlices(expected, actual) {
		t.Errorf("Partition failed. Expected: %v, Got: %v", expected, actual)
	}
}

// Helper function to compare two slices
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
