package stack_min

import (
	"log"
	"math"
	"testing"
)

// Logger function to log test results
func logTestResult(t *testing.T, message string, args ...interface{}) {
	log.Printf(message, args...)
	if t.Failed() {
		t.Fail()
	}
}

// TestMinStack checks basic operations and minimum value tracking
func TestMinStack(t *testing.T) {
	stack := NewMinStack()

	// Test initial min value
	logTestResult(t, "Initial Min() = %v; expected math.MaxInt", stack.Min())
	if got := stack.Min(); got != math.MaxInt {
		t.Errorf("Min() = %v; want %v", got, math.MaxInt)
	}

	// Push elements and check minimum
	stack.Push(10)
	logTestResult(t, "After pushing 10, Min() = %v; expected 10", stack.Min())
	if got := stack.Min(); got != 10 {
		t.Errorf("Min() = %v; want 10", got)
	}

	stack.Push(5)
	logTestResult(t, "After pushing 5, Min() = %v; expected 5", stack.Min())
	if got := stack.Min(); got != 5 {
		t.Errorf("Min() = %v; want 5", got)
	}

	stack.Push(15)
	logTestResult(t, "After pushing 15, Min() = %v; expected 5", stack.Min())
	if got := stack.Min(); got != 5 {
		t.Errorf("Min() = %v; want 5", got)
	}

	// Test popping elements
	stack.Pop()
	logTestResult(t, "After popping 15, Min() = %v; expected 5", stack.Min())
	if got := stack.Min(); got != 5 {
		t.Errorf("Min() = %v; want 5", got)
	}

	stack.Pop()
	logTestResult(t, "After popping 5, Min() = %v; expected 10", stack.Min())
	if got := stack.Min(); got != 10 {
		t.Errorf("Min() = %v; want 10", got)
	}

	stack.Pop()
	logTestResult(t, "After popping 10, Min() = %v; expected math.MaxInt", stack.Min())
	if got := stack.Min(); got != math.MaxInt {
		t.Errorf("Min() = %v; want %v", got, math.MaxInt)
	}
}

// TestMinStack_EmptyStack checks behavior of an empty stack
func TestMinStack_EmptyStack(t *testing.T) {
	stack := NewMinStack()

	// Check min on an empty stack
	logTestResult(t, "After popping from empty stack, Min() = %v; expected math.MaxInt", stack.Min())
	if got := stack.Min(); got != math.MaxInt {
		t.Errorf("Min() = %v; want %v", got, math.MaxInt)
	}

	stack.Pop() // Ensure that popping from empty stack does not crash
	logTestResult(t, "After popping from empty stack, Min() = %v; expected math.MaxInt", stack.Min())
	if got := stack.Min(); got != math.MaxInt {
		t.Errorf("Min() = %v; want %v", got, math.MaxInt)
	}
}

// TestMinStack_PushPopEdgeCases checks edge cases with repeated and extreme values
func TestMinStack_PushPopEdgeCases(t *testing.T) {
	stack := NewMinStack()

	// Push elements with extreme values
	stack.Push(math.MaxInt - 1)
	stack.Push(math.MaxInt - 2)
	stack.Push(math.MinInt)
	stack.Push(0)
	stack.Push(math.MaxInt)

	logTestResult(t, "After pushing extreme values, Min() = %v; expected %v", stack.Min(), math.MinInt)
	if got := stack.Min(); got != math.MinInt {
		t.Errorf("Min() = %v; want %v", got, math.MinInt)
	}

	// Pop elements and check min
	stack.Pop()
	logTestResult(t, "After popping math.MaxInt, Min() = %v; expected %v", stack.Min(), math.MinInt)
	if got := stack.Min(); got != math.MinInt {
		t.Errorf("Min() = %v; want %v", got, math.MinInt)
	}

	stack.Pop()
	logTestResult(t, "After popping 0, Min() = %v; expected %v", stack.Min(), math.MinInt)
	if got := stack.Min(); got != math.MinInt {
		t.Errorf("Min() = %v; want %v", got, math.MinInt)
	}

	stack.Pop()
	logTestResult(t, "After popping math.MinInt, Min() = %v; expected %v", stack.Min(), math.MaxInt-2)
	if got := stack.Min(); got != math.MaxInt-2 {
		t.Errorf("Min() = %v; want %v", got, math.MaxInt-2)
	}

	stack.Pop()
	logTestResult(t, "After popping math.MaxInt-2, Min() = %v; expected %v", stack.Min(), math.MaxInt-1)
	if got := stack.Min(); got != math.MaxInt-1 {
		t.Errorf("Min() = %v; want %v", got, math.MaxInt-1)
	}

	stack.Pop()
	logTestResult(t, "After popping math.MaxInt-1, Min() = %v; expected math.MaxInt", stack.Min())
	if got := stack.Min(); got != math.MaxInt {
		t.Errorf("Min() = %v; want %v", got, math.MaxInt)
	}
}
