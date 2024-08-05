package queue_via_stacks

import (
	"testing"
)

// TestPushAndPop tests basic push and pop operations
func TestPushAndPop(t *testing.T) {
	q := NewMyQueue()

	// Test pushing elements
	q.Push(1)
	q.Push(2)
	q.Push(3)
	t.Logf("Pushed elements: 1, 2, 3")
	// Test popping elements
	if got := q.Pop(); got != 1 {
		t.Errorf("Pop() = %d; want 1", got)
	} else {
		t.Logf("Popped element: 1")
	}
	if got := q.Pop(); got != 2 {
		t.Errorf("Pop() = %d; want 2", got)
	} else {
		t.Logf("Popped element: 2")
	}
	if got := q.Pop(); got != 3 {
		t.Errorf("Pop() = %d; want 3", got)
	} else {
		t.Logf("Popped element: 3")
	}
}

// TestPushAndPeek tests pushing and peeking operations
func TestPushAndPeek(t *testing.T) {
	q := NewMyQueue()

	q.Push(1)
	q.Push(2)
	t.Logf("Pushed elements: 1, 2")
	// Test peeking elements
	if got := q.Peek(); got != 1 {
		t.Errorf("Peek() = %d; want 1", got)
	} else {
		t.Logf("Peeked element: 1")
	}

	q.Push(3)
	t.Logf("Pushed element: 3")
	if got := q.Peek(); got != 1 {
		t.Errorf("Peek() = %d; want 1", got)
	} else {
		t.Logf("Peeked element: 1")
	}
}

// TestPopEmptyQueue tests popping from an empty queue
func TestPopEmptyQueue(t *testing.T) {
	q := NewMyQueue()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when popping from empty queue")
		}
	}()

	q.Pop() // This should panic
}

// TestPeekEmptyQueue tests peeking from an empty queue
func TestPeekEmptyQueue(t *testing.T) {
	q := NewMyQueue()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when peeking from empty queue")
		}
	}()

	q.Peek() // This should panic
}

// TestEmpty checks the empty status of the queue
func TestEmpty(t *testing.T) {
	q := NewMyQueue()

	t.Logf("Queue is empty: %t", q.Empty())
	if !q.Empty() {
		t.Errorf("Empty() = false; want true")
	}

	q.Push(1)
	t.Logf("Pushed element: 1. Queue is not empty: %t", q.Empty())
	if q.Empty() {
		t.Errorf("Empty() = true; want false")
	} else {
		t.Logf("Queue is not empty: %t", q.Empty())
	}

	p := q.Pop()
	t.Logf("Popped element: %d. Queue is empty: %t", p, q.Empty())
	if !q.Empty() {
		t.Errorf("Empty() = false; want true")
	} else {
		t.Logf("Queue is empty: %t", q.Empty())
	}
}

// TestOrderAfterPushPop tests the order of elements after various push and pop operations
func TestOrderAfterPushPop(t *testing.T) {
	q := NewMyQueue()

	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	t.Logf("Pushed elements: 1, 2, 3, 4")

	if got := q.Pop(); got != 1 {
		t.Errorf("Pop() = %d; want 1", got)
	} else {
		t.Logf("Popped element: 1")
	}
	if got := q.Peek(); got != 2 {
		t.Errorf("Peek() = %d; want 2", got)
	} else {
		t.Logf("Peeked element: 2")
	}
	q.Push(5)
	t.Logf("Pushed element: 5")
	if got := q.Pop(); got != 2 {
		t.Errorf("Pop() = %d; want 2", got)
	} else {
		t.Logf("Popped element: 2")
	}
	if got := q.Pop(); got != 3 {
		t.Errorf("Pop() = %d; want 3", got)
	} else {
		t.Logf("Popped element: 3")
	}
	if got := q.Pop(); got != 4 {
		t.Errorf("Pop() = %d; want 4", got)
	} else {
		t.Logf("Popped element: 4")
	}
	if got := q.Pop(); got != 5 {
		t.Errorf("Pop() = %d; want 5", got)
	} else {
		t.Logf("Popped element: 5")
	}
}
