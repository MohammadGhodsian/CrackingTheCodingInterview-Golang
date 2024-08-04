package return_kth_to_last

import "CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"

/*
2.2 Return Kth to Last
Implement an algorithm to find the kth to last element of a singly linked list.
*/

func ReturnKthToLast(node *LinkedList.ListNode, k int) (*LinkedList.ListNode, int) {
	if node == nil || node.Next == nil {
		return node, 0
	}
	n, i := ReturnKthToLast(node.Next, k)
	i++
	if i == k {
		return node, i
	}
	if i < k {
		return nil, i
	}
	return n, i
}
