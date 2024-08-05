package loop_detection

import (
	"CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"
	"testing"
)

// Helper function to create a list with a loop for testing
func createListWithLoop(values []int, loopStartIndex int) *LinkedList.ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &LinkedList.ListNode{Value: values[0]}
	current := head
	var loopStart *LinkedList.ListNode

	for i := 1; i < len(values); i++ {
		current.Next = &LinkedList.ListNode{Value: values[i]}
		current = current.Next
		if i == loopStartIndex {
			loopStart = current
		}
	}

	if loopStart != nil {
		current.Next = loopStart
	}

	return head
}

// Test case 1: List with a loop
func TestLoopDetection_WithLoop(t *testing.T) {
	head := createListWithLoop([]int{1, 2, 3, 4, 5}, 2) // Loop starts at node with value 3
	expected := head.Next.Next                          // Node with value 3

	detected := LoopDetection(head)
	if detected != expected {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected, detected)
	} else {
		t.Logf("Test case 1 succeeded. List: %v, expected: %v, got: %v", head, expected, detected)
	}
}

// Test case 2: List without a loop
func TestLoopDetection_WithoutLoop(t *testing.T) {
	head := LinkedList.CreateList([]int{1, 2, 3, 4, 5})

	detected := LoopDetection(head)
	if detected != nil {
		t.Errorf("Test case 2 failed: expected nil, got %v", detected)
	} else {
		t.Logf("Test case 2 succeeded. List: %v, expected: nil, got: %v", head, detected)
	}
}

// Test case 3: Empty list
func TestLoopDetection_EmptyList(t *testing.T) {
	var head *LinkedList.ListNode

	detected := LoopDetection(head)
	if detected != nil {
		t.Errorf("Test case 3 failed: expected nil, got %v", detected)
	} else {
		t.Logf("Test case 3 succeeded. List: %v, expected: nil, got: %v", head, detected)
	}
}

// Test case 4: List with single node without loop
func TestLoopDetection_SingleNode_NoLoop(t *testing.T) {
	head := &LinkedList.ListNode{Value: 1}

	detected := LoopDetection(head)
	if detected != nil {
		t.Errorf("Test case 4 failed: expected nil, got %v", detected)
	} else {
		t.Logf("Test case 4 succeeded. List: %v, expected: nil, got: %v", head, detected)
	}
}

// Test case 5: List with single node with a loop
func TestLoopDetection_SingleNode_WithLoop(t *testing.T) {
	head := &LinkedList.ListNode{Value: 1}
	head.Next = head // Loop to itself

	detected := LoopDetection(head)
	if detected != head {
		t.Errorf("Test case 5 failed: expected %v, got %v", head, detected)
	} else {
		t.Logf("Test case 5 succeeded. List: %v, expected: %v, got: %v", head, head, detected)
	}
}
