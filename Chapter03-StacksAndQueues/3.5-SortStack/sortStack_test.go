package sort_stack

import (
	"fmt"
	"testing"
)

// Helper function to convert stack to a slice for easier verification
func stackToSlice(stack *Stack) []int {
	slice := []int{}
	for !stack.IsEmpty() {
		slice = append(slice, stack.Pop())
	}
	// Reverse the slice to match the original order
	for i := len(slice) - 1; i >= 0; i-- {
		stack.Push(slice[i])
	}
	return slice
}

// Log the state of the stack
func logStack(stack *Stack, message string) {
	fmt.Println(message)
	tempStack := NewStack()
	for !stack.IsEmpty() {
		item := stack.Pop()
		fmt.Printf("%d ", item)
		tempStack.Push(item)
	}
	fmt.Println()
	// Restore the stack
	for !tempStack.IsEmpty() {
		stack.Push(tempStack.Pop())
	}
}

// TestSortStackMultipleElements tests sorting a stack with multiple elements
func TestSortStackMultipleElements(t *testing.T) {
	stack := NewStack()
	stack.Push(3)
	stack.Push(1)
	stack.Push(4)
	stack.Push(2)

	fmt.Println("Initial stack state:")
	logStack(stack, "Before sorting:")

	SortStack(stack)

	fmt.Println("Stack state after sorting:")
	logStack(stack, "After sorting:")

	expected := []int{1, 2, 3, 4}
	got := stackToSlice(stack)
	if len(got) != len(expected) {
		t.Errorf("TestSortStackMultipleElements failed: expected %v, got %v", expected, got)
	}
	for i, v := range got {
		if v != expected[i] {
			t.Errorf("TestSortStackMultipleElements failed: expected %v, got %v", expected, got)
			return
		}
	}
}

// TestSortStackSingleElement tests sorting a stack with a single element
func TestSortStackSingleElement(t *testing.T) {
	stack := NewStack()
	stack.Push(42)

	fmt.Println("Initial stack state:")
	logStack(stack, "Before sorting:")

	SortStack(stack)

	fmt.Println("Stack state after sorting:")
	logStack(stack, "After sorting:")

	expected := []int{42}
	got := stackToSlice(stack)
	if len(got) != len(expected) {
		t.Errorf("TestSortStackSingleElement failed: expected %v, got %v", expected, got)
	}
	for i, v := range got {
		if v != expected[i] {
			t.Errorf("TestSortStackSingleElement failed: expected %v, got %v", expected, got)
			return
		}
	}
}

// TestSortStackAlreadySorted tests sorting an already sorted stack
func TestSortStackAlreadySorted(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println("Initial stack state:")
	logStack(stack, "Before sorting:")

	SortStack(stack)

	fmt.Println("Stack state after sorting:")
	logStack(stack, "After sorting:")

	expected := []int{1, 2, 3, 4}
	got := stackToSlice(stack)
	if len(got) != len(expected) {
		t.Errorf("TestSortStackAlreadySorted failed: expected %v, got %v", expected, got)
	}
	for i, v := range got {
		if v != expected[i] {
			t.Errorf("TestSortStackAlreadySorted failed: expected %v, got %v", expected, got)
			return
		}
	}
}

// TestSortStackEmpty tests sorting an empty stack
func TestSortStackEmpty(t *testing.T) {
	stack := NewStack()

	fmt.Println("Initial stack state:")
	logStack(stack, "Before sorting:")

	SortStack(stack)

	fmt.Println("Stack state after sorting:")
	logStack(stack, "After sorting:")

	expected := []int{}
	got := stackToSlice(stack)
	if len(got) != len(expected) {
		t.Errorf("TestSortStackEmpty failed: expected %v, got %v", expected, got)
	}
}
