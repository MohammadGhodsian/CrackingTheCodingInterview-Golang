package loop_detection

import "CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"

/*
2.8 Loop Detection
Given a circular linked list, implement an algorithm that returns the node at the beginning of the loop.
DEFINITION
Circular linked list: A (corrupt) linked list in which a node's next pointer points to an earlier node,
 so as to make a loop in the linked list.
EXAMPLE
Input: A -> B -> C -> D -> E -> C (the same C as earlier)
Output: C
*/

func LoopDetection(head *LinkedList.ListNode) *LinkedList.ListNode {
	if head == nil {
		return nil
	}
	// Using Floyd’s Cycle-Finding Algorithm (also known as the Tortoise and Hare algorithm)
	slow := head
	fast := head

	// Assume there is a loop in the linked list.
	// As fast moves twice as fast, it will enter the loop and start circling around it.
	// Eventually, it will "lap" the slow pointer, which is moving slower and also in the loop.

	// Phase 1: Detect if a loop exists
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	// No loop detected
	if fast == nil || fast.Next == nil {
		return nil
	}
	// Loop detected

	// Phase 2: Find the start node of the loop
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
