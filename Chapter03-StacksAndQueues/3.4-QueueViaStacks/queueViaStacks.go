package queue_via_stacks

/*
3.4 Queue via Stacks
Implement a MyQueue class which implements a queue using two stacks.
*/

// MyQueue represents a queue implemented using two stacks
type MyQueue struct {
	stack1 *Stack
	stack2 *Stack
}

// NewMyQueue creates a new instance of MyQueue
func NewMyQueue() *MyQueue {
	return &MyQueue{
		stack1: NewStack(),
		stack2: NewStack(),
	}
}

// Push adds an element to the end of the queue
func (q *MyQueue) Push(x int) {
	q.stack1.Push(x)
}

// Pop removes the element from the front of the queue
func (q *MyQueue) Pop() int {
	if q.stack2.IsEmpty() {
		// Transfer all elements from stack1 to stack2
		for !q.stack1.IsEmpty() {
			q.stack2.Push(q.stack1.Pop())
		}
	}
	if q.stack2.IsEmpty() {
		panic("pop from empty queue")
	}
	// Pop the top element from stack2
	return q.stack2.Pop()
}

// Peek returns the element at the front of the queue without removing it
func (q *MyQueue) Peek() int {
	if q.stack2.IsEmpty() {
		// Transfer all elements from stack1 to stack2
		for !q.stack1.IsEmpty() {
			q.stack2.Push(q.stack1.Pop())
		}
	}
	if q.stack2.IsEmpty() {
		panic("peek from empty queue")
	}
	// Peek the top element from stack2
	return q.stack2.Peek()
}

// Empty returns whether the queue is empty
func (q *MyQueue) Empty() bool {
	return q.stack1.IsEmpty() && q.stack2.IsEmpty()
}

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

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.items)
}
