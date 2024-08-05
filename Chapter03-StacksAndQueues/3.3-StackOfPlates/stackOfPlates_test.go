package stack_of_plates

import (
	"fmt"
	"testing"
)

// TestPushAndPop tests basic push and pop operations
func TestPushAndPop(t *testing.T) {
	stack := NewSetOfStacks(3)
	fmt.Println("Running TestPushAndPop")

	// Test pushing elements
	fmt.Println("Pushing elements: 1, 2, 3")
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Test popping elements
	fmt.Println("Popping elements")
	if got := stack.Pop(); got != 3 {
		t.Errorf("Pop() = %d; want 3", got)
	}
	if got := stack.Pop(); got != 2 {
		t.Errorf("Pop() = %d; want 2", got)
	}
	if got := stack.Pop(); got != 1 {
		t.Errorf("Pop() = %d; want 1", got)
	}
}

// TestPushToNewStack tests pushing elements that create new stacks
func TestPushToNewStack(t *testing.T) {
	stack := NewSetOfStacks(2)
	fmt.Println("Running TestPushToNewStack")

	stack.Push(1)
	stack.Push(2) // Should be in stack 0
	fmt.Println("Pushed 1 and 2 to stack 0")

	// This push should cause the creation of a new stack
	fmt.Println("Pushing 3, which should create a new stack")
	stack.Push(3)

	fmt.Println("Popping elements")
	if got := stack.Pop(); got != 3 {
		t.Errorf("Pop() = %d; want 3", got)
	}
	if got := stack.Pop(); got != 2 {
		t.Errorf("Pop() = %d; want 2", got)
	}
	if got := stack.Pop(); got != 1 {
		t.Errorf("Pop() = %d; want 1", got)
	}
}

// TestPopEmptyStack tests popping from an empty SetOfStacks
func TestPopEmptyStack(t *testing.T) {
	stack := NewSetOfStacks(3)
	fmt.Println("Running TestPopEmptyStack")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when popping from empty stack")
		}
	}()

	fmt.Println("Attempting to pop from an empty stack")
	stack.Pop() // This should panic
}

// TestPopAt tests the popAt function
func TestPopAt(t *testing.T) {
	stack := NewSetOfStacks(2)
	fmt.Println("Running TestPopAt")

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println("Pushed 1, 2 to stack 0 and 3, 4 to stack 1")
	fmt.Println("Performing PopAt(0)")

	if got := stack.PopAt(0); got != 2 {
		t.Errorf("PopAt(0) = %d; want 2", got)
	}

	fmt.Println("Popping remaining elements")
	if got := stack.Pop(); got != 4 {
		t.Errorf("Pop() = %d; want 4", got)
	}
	if got := stack.Pop(); got != 3 {
		t.Errorf("Pop() = %d; want 3", got)
	}

	// Trying to pop at an index that doesn't exist
	fmt.Println("Attempting to popAt(1) from a single stack")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when popping at invalid index")
		}
	}()
	stack.PopAt(1) // This should panic
}
