package sort_stack

// Stack represents a stack data structure
type Stack struct {
	items []int
}

// NewStack creates a new stack
func NewStack() *Stack {
	return &Stack{}
}

// Push adds an element to the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("pop from empty stack")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("peek from empty stack")
	}
	return s.items[len(s.items)-1]
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// SortStack sorts the stack such that the smallest items are on top
func SortStack(original *Stack) {
	tempStack := NewStack()

	for !original.IsEmpty() {
		// Pop the top element from the original stack
		temp := original.Pop()

		// While the temporary stack is not empty and top of tempStack is greater than temp
		for !tempStack.IsEmpty() && tempStack.Peek() > temp {
			// Move elements from tempStack to original
			original.Push(tempStack.Pop())
		}

		// Push temp onto tempStack
		tempStack.Push(temp)
	}

	// Move elements back to the original stack
	for !tempStack.IsEmpty() {
		original.Push(tempStack.Pop())
	}
}
