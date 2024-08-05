package palindrome

/*
2.6 Palindrome
Implement a function to check if a linked list is a palindrome.
*/

import "CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"

func IsPalindrome(head *LinkedList.ListNode) bool {
	var headOfReversedList *LinkedList.ListNode
	for node := head; node != nil; node = node.Next {
		tempNode := LinkedList.ListNode{Value: node.Value}
		if headOfReversedList != nil {
			tempNode.Next = headOfReversedList
		}
		headOfReversedList = &tempNode
	}
	return LinkedList.ListsEqual(head, headOfReversedList)
}
