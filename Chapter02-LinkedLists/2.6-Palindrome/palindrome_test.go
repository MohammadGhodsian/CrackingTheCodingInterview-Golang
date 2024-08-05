package palindrome

import (
	"CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{input: []int{1, 2, 3, 2, 1}, expected: true},
		{input: []int{1, 2, -1, 0, -1, 2, 1}, expected: true},
		{input: []int{1, 2, 3, 4, 5}, expected: false},
		{input: []int{1}, expected: true},
		{input: []int{1, 2, 3, 4, 3, 2, 1}, expected: true},
		{input: []int{}, expected: true},
	}

	for _, test := range tests {
		head := LinkedList.CreateList(test.input)
		result := IsPalindrome(head)
		if result != test.expected {
			t.Errorf("IsPalindrome(%v) = %v; want %v", test.input, result, test.expected)
		} else {
			t.Logf("Success: IsPalindrome(%v) = %v", test.input, result)
		}
	}
}
