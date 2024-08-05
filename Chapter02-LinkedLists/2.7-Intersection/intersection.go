package intersection

import "CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"

/*
2.7 Intersection
Given two (singly) linked lists, determine if the two lists intersect.
Return the intersecting node.
Note that the intersection is defined based on reference, not value.
That is, if the kth node of the first linked list is the exact same node (by reference) as the jth node of the second linked list, then they are intersecting.
*/

func FindIntersection(headA, headB *LinkedList.ListNode) (intersection *LinkedList.ListNode) {
	if headA == nil || headB == nil {
		return nil
	}
	// Get the length of each linked list
	lenA := headA.Length()
	lenB := headB.Length()
	// Align the start points (If there is an intersection, the remain list length should be the same)
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	}
	// Traverse both lists
	for headA != nil && headB != nil {
		if headA == headB {
			intersection = headA
			break
		}
		headA = headA.Next
		headB = headB.Next
	}
	return
}
