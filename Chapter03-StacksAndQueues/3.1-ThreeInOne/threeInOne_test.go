package three_in_one

import (
	"testing"
)

func TestThreeInOne(t *testing.T) {
	totalSize := 10
	threeStacks := NewThreeInOne(totalSize)

	// Test pushing values into each stack
	tests := []struct {
		stackNum int
		value    int
	}{
		{0, 1},
		{1, 2},
		{2, 3},
		{0, 4},
		{1, 5},
		{2, 6},
	}

	for _, test := range tests {
		err := threeStacks.Push(test.value, test.stackNum)
		if err != nil {
			t.Errorf("Push(%d, %d) failed: %v", test.value, test.stackNum, err)
		}
	}

	// Test peeking values from each stack
	peekTests := []struct {
		stackNum int
		expected int
	}{
		{0, 4},
		{1, 5},
		{2, 6},
	}

	for _, test := range peekTests {
		value, err := threeStacks.Peek(test.stackNum)
		if err != nil {
			t.Errorf("Peek(%d) failed: %v", test.stackNum, err)
		} else if value != test.expected {
			t.Errorf("Peek(%d) = %d; want %d", test.stackNum, value, test.expected)
		}
	}

	// Test popping values from each stack
	popTests := []struct {
		stackNum int
		expected int
	}{
		{0, 4},
		{1, 5},
		{2, 6},
	}

	for _, test := range popTests {
		value, err := threeStacks.Pop(test.stackNum)
		if err != nil {
			t.Errorf("Pop(%d) failed: %v", test.stackNum, err)
		} else if value != test.expected {
			t.Errorf("Pop(%d) = %d; want %d", test.stackNum, value, test.expected)
		}
	}

	// Test checking if stacks are empty
	emptyTests := []struct {
		stackNum int
		expected bool
	}{
		{0, false},
		{1, false},
		{2, false},
	}

	for _, test := range emptyTests {
		isEmpty, err := threeStacks.IsEmpty(test.stackNum)
		if err != nil {
			t.Errorf("IsEmpty(%d) failed: %v", test.stackNum, err)
		} else if isEmpty != test.expected {
			t.Errorf("IsEmpty(%d) = %v; want %v", test.stackNum, isEmpty, test.expected)
		}
	}

	// Test popping all values and checking if stacks are empty
	for i := 0; i < numStacks; i++ {
		_, err := threeStacks.Pop(i)
		if err != nil {
			t.Errorf("Pop(%d) failed: %v", i, err)
		}
	}

	for i := 0; i < numStacks; i++ {
		isEmpty, err := threeStacks.IsEmpty(i)
		if err != nil {
			t.Errorf("IsEmpty(%d) failed: %v", i, err)
		} else if !isEmpty {
			t.Errorf("IsEmpty(%d) = %v; want true", i, isEmpty)
		}
	}

	// Test pushing beyond capacity
	err := threeStacks.Push(100, 0)
	if err != nil {
		t.Errorf("Push(100, 0) = %v; err", err)
	}
	err = threeStacks.Push(200, 0)
	if err != nil {
		t.Errorf("Push(200, 0) = %v; err", err)
	}
	err = threeStacks.Push(300, 0)
	if err != nil {
		t.Errorf("Push(300, 0) = %v; err", err)
	}
	err = threeStacks.Push(400, 0)
	if err != ErrStackFull {
		t.Errorf("Push(400, 0) = %v; want ErrStackFull", err)
	}
}
