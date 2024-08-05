package three_in_one

import (
	"errors"
	"fmt"
)

/*
3.1 Three in One
Describe how you could use a single array to implement three stacks.
*/

type stack struct {
	size       int
	startIndex int
	freeIndex  int
}

type ThreeInOne struct {
	List []int
	s    [numStacks]stack
}

const numStacks = 3

var ErrInvalidStackNumber error = errors.New("invalid stack number")
var ErrStackFull error = errors.New("stack is full")
var ErrStackEmpty error = errors.New("stack is empty")

// NewThreeInOne initializes a ThreeInOne with a total size
func NewThreeInOne(totalSize int) *ThreeInOne {
	if totalSize < numStacks {
		totalSize = numStacks
	}
	stackSize := totalSize / numStacks

	return &ThreeInOne{
		List: make([]int, totalSize),
		s: [numStacks]stack{
			{stackSize, 0, 0},
			{stackSize, stackSize, stackSize},
			{totalSize - 2*stackSize, 2 * stackSize, 2 * stackSize},
		},
	}
}

// Push adds a value to the specified stack
func (t *ThreeInOne) Push(value int, stackNum int) error {
	if stackNum < 0 || stackNum >= numStacks {
		return ErrInvalidStackNumber
	}
	stk := &t.s[stackNum]
	fmt.Printf("Attempting to push to stack %d. Current freeIndex: %d, startIndex: %d, stack size: %d\n", stackNum, stk.freeIndex, stk.startIndex, stk.size)
	if stk.freeIndex >= stk.startIndex+stk.size {
		fmt.Printf("Push failed for stack %d: Stack is full\n", stackNum)
		return ErrStackFull
	}
	t.List[stk.freeIndex] = value
	stk.freeIndex++
	fmt.Printf("Pushed %d to stack %d. New freeIndex: %d\n", value, stackNum, stk.freeIndex)
	return nil
}

// Pop removes and returns the top value from the specified stack
func (t *ThreeInOne) Pop(stackNum int) (int, error) {
	if stackNum < 0 || stackNum >= numStacks {
		return 0, ErrInvalidStackNumber
	}
	stk := &t.s[stackNum]
	if stk.freeIndex == stk.startIndex {
		return 0, ErrStackEmpty
	}
	stk.freeIndex--
	value := t.List[stk.freeIndex]
	fmt.Printf("Popped %d from stack %d. New freeIndex: %d\n", value, stackNum, stk.freeIndex)
	return value, nil
}

// Peek returns the top value of the specified stack without removing it
func (t *ThreeInOne) Peek(stackNum int) (int, error) {
	if stackNum < 0 || stackNum >= numStacks {
		return 0, ErrInvalidStackNumber
	}
	stk := &t.s[stackNum]
	if stk.freeIndex == stk.startIndex {
		return 0, ErrStackEmpty
	}
	value := t.List[stk.freeIndex-1]
	fmt.Printf("Peeked at stack %d: %d\n", stackNum, value)
	return value, nil
}

// IsEmpty checks if the specified stack is empty
func (t *ThreeInOne) IsEmpty(stackNum int) (bool, error) {
	if stackNum < 0 || stackNum >= numStacks {
		return false, ErrInvalidStackNumber
	}
	stk := &t.s[stackNum]
	isEmpty := stk.freeIndex == stk.startIndex
	fmt.Printf("Is stack %d empty? %v\n", stackNum, isEmpty)
	return isEmpty, nil
}
