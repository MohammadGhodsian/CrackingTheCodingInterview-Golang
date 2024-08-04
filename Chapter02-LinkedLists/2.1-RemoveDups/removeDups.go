package remove_dups

import (
	linkedList "CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"
)

/*
2.1 Remove Dups
Write code to remove duplicates from an unsorted linked list.
How would you solve this problem if a temporary buffer is not allowed?
*/

func RemoveDuplicatesNoBuffer(head *linkedList.ListNode) *linkedList.ListNode {
	// Start with the head of the list.
	current := head

	// Traverse the list with the `current` node.
	for current != nil {
		// For each node in the list, use `inner` to check the rest of the list for duplicates.
		inner := current

		// Traverse the rest of the list starting from the node after `current`.
		for inner.Next != nil {
			// If a duplicate value is found (value of `inner.Next` is the same as `current`), remove it.
			if inner.Next.Value == current.Value {
				// Skip the next node by pointing `inner.Next` to `inner.Next.Next`.
				inner.Next = inner.Next.Next
			} else {
				// Otherwise, move `inner` to the next node.
				inner = inner.Next
			}
		}
		// Move `current` to the next node in the list.
		current = current.Next
	}

	// Return the modified head of the list.
	return head
}
