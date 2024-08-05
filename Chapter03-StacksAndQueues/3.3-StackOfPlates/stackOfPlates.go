package stack_of_plates

/*
3.3 Stack of Plates
Imagine a (literal) stack of plates.
If the stack gets too high, it might topple.
Therefore, in real life, we would likely start a new stack when the previous stack exceeds some threshold.
Implement a data structure SetOfStacks that mimics this.
SetOfStacks should be composed of several stacks and should create a new stack once the previous one exceeds capacity.
SetOfStacks.push() and SetOfStacks.pop() should behave identically to a single stack (that is, pop() should return the same values as it would if there were just a single stack).
FOLLOW UP
Implement a function popAt(int index) which performs a pop operation on a specific sub-stack.
*/

// SetOfStacks represents a collection of stacks
type SetOfStacks struct {
	stacks   [][]int
	capacity int
}

// NewSetOfStacks creates a new SetOfStacks with the given stack capacity
func NewSetOfStacks(capacity int) *SetOfStacks {
	return &SetOfStacks{
		stacks:   [][]int{},
		capacity: capacity,
	}
}

// Push adds a value to the SetOfStacks
func (s *SetOfStacks) Push(value int) {
	if len(s.stacks) == 0 || len(s.stacks[len(s.stacks)-1]) >= s.capacity {
		s.stacks = append(s.stacks, []int{})
	}
	s.stacks[len(s.stacks)-1] = append(s.stacks[len(s.stacks)-1], value)
}

// Pop removes and returns the top value from the SetOfStacks
func (s *SetOfStacks) Pop() int {
	if len(s.stacks) == 0 {
		panic("pop from empty stack")
	}
	stack := s.stacks[len(s.stacks)-1]
	value := stack[len(stack)-1]
	s.stacks[len(s.stacks)-1] = stack[:len(stack)-1]
	if len(s.stacks[len(s.stacks)-1]) == 0 {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}
	return value
}

// PopAt pops a value from a specific stack in the SetOfStacks
func (s *SetOfStacks) PopAt(index int) int {
	if index < 0 || index >= len(s.stacks) {
		panic("index out of range")
	}
	stack := s.stacks[index]
	if len(stack) == 0 {
		panic("pop from empty stack")
	}
	value := stack[len(stack)-1]
	s.stacks[index] = stack[:len(stack)-1]

	// If the stack at index becomes empty, remove it
	if len(s.stacks[index]) == 0 {
		s.stacks = append(s.stacks[:index], s.stacks[index+1:]...)
	}

	return value
}
