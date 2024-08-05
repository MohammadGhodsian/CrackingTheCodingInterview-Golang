package stack_min

import (
	"math"
)

// MinStack defines the stack with min functionality
type MinStack struct {
	stack    []int
	minStack []int
}

// NewMinStack creates a new MinStack
func NewMinStack() *MinStack {
	return &MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt},
	}
}

// Push adds an element to the stack
func (ms *MinStack) Push(value int) {
	ms.stack = append(ms.stack, value)
	if value <= ms.minStack[len(ms.minStack)-1] {
		ms.minStack = append(ms.minStack, value)
	}
}

// Pop removes the top element from the stack
func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}
	value := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	if value == ms.minStack[len(ms.minStack)-1] {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}
}

// Min returns the minimum element in the stack
func (ms *MinStack) Min() int {
	if len(ms.minStack) == 0 {
		return math.MaxInt
	}
	return ms.minStack[len(ms.minStack)-1]
}
