package sum_lists

/*
Sum Lists
You have two numbers represented by a linked list, where each node contains a single digit. The digits are stored in reverse order, such that the 1's digit is at the head of the list. Write a function that adds the two numbers and returns the sum as a linked list.
EXAMPLE
Input: (7->1->6) + (5->9->2).Thatis, 617 + 295.
Output: 2->1->9.That is, 912.
Suppose the digits are stored in forward order. Repeat the above problem.
EXAMPLE
Input: (6->1->7) + (2->9->5). That is, 617 + 295.
Output: 9->1->2. That is, 912.
*/

import (
	"CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"
	"fmt"
	"math"
)

func SumListsReverseOrder(l1, l2 *LinkedList.ListNode) *LinkedList.ListNode {
	num1 := listToNumberReverse(l1)
	num2 := listToNumberReverse(l2)

	sum := num1 + num2

	fmt.Printf(" SumListsReverseOrder ***** num1: %d, num2: %d, sum: %d\n", num1, num2, sum)
	return numberToListReverse(sum)
}

// Helper function to convert linked list to number
func listToNumberReverse(head *LinkedList.ListNode) int {
	num := 0
	factor := 1
	for head != nil {
		num += head.Value * factor
		factor *= 10
		head = head.Next
	}
	return num
}

// Helper function to convert number to linked list
// Helper function to convert number to linked list in reverse order
func numberToListReverse(num int) *LinkedList.ListNode {
	if num == 0 {
		return &LinkedList.ListNode{Value: 0}
	}

	var head *LinkedList.ListNode
	var current *LinkedList.ListNode
	for num > 0 {
		newNode := &LinkedList.ListNode{Value: num % 10}
		if head == nil {
			head = newNode
			current = head
		} else {
			current.Next = newNode
			current = current.Next
		}
		num /= 10
	}
	return head
}

func SumListsForwardOrder(l1, l2 *LinkedList.ListNode) *LinkedList.ListNode {
	num1 := listToNumberForward(l1)
	num2 := listToNumberForward(l2)
	sum := num1 + num2
	fmt.Printf(" SumListsReverseOrder ***** num1: %d, num2: %d, sum: %d\n", num1, num2, sum)

	return numberToListForward(sum)
}

// Helper function to convert linked list to number (forward order)
func listToNumberForward(head *LinkedList.ListNode) int {
	num := 0
	// First find the length of the list to determine the factor
	length := head.Length()
	for head != nil {
		num += head.Value * int(math.Pow(10, float64(length-1)))
		head = head.Next
		length--
	}
	return num
}

// Helper function to convert number to linked list (forward order)
func numberToListForward(num int) *LinkedList.ListNode {
	if num == 0 {
		return &LinkedList.ListNode{Value: 0}
	}

	// Create a list of digits from the number
	var digits []int
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	// Reverse the slice to get the correct forward order
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	// Create the linked list from the digits in forward order
	var head, current *LinkedList.ListNode
	for _, digit := range digits {
		newNode := &LinkedList.ListNode{Value: digit}
		if head == nil {
			head = newNode
			current = head
		} else {
			current.Next = newNode
			current = newNode
		}
	}
	return head
}
