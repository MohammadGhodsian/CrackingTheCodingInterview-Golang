package delete_middle_node

/*
Delete Middle Node:
Implement an algorithm to delete a node in the middle (i.e., any node but the first and last node, not necessarily the exact middle) of a singly linked list, given only access to that node.
EXAMPLE
Input: the node c from the linked list a->b->c->d->e->f
Result: nothing is returned, but the new linked list looks like a - >b- >d - >e- >f
*/

import "CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"

func DeleteMiddleNode(node *LinkedList.ListNode) {
	if node == nil || node.Next == nil {
		// Invalid case: node is null or the node is the last node
		return
	}
	// Copy the value of the next node to the current node
	node.Value = node.Next.Value
	// Bypass the next node
	node.Next = node.Next.Next
}
