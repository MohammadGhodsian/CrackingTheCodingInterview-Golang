package partition

/*
2.4 Partition
Write code to partition a linked list around a value x,
such that all nodes less than x come before all nodes greater than or equal to x.
lf x is contained within the list,
the values of x only need to be after the elements less than x (see below).
The partition element x can appear anywhere in the "right partition";
it does not need to appear between the left and right partitions.

EXAMPLE
Input: 3->5->8->5->10->2->1 (partition=5)
Output: 3->1->2->10->5->5->8
*/

import "CrackingTheCodingInterview/Chapter02-LinkedLists/LinkedList"

func Partition(head *LinkedList.ListNode, x int) *LinkedList.ListNode {
	if head == nil {
		return nil
	}

	lessHead := &LinkedList.ListNode{}
	moreHead := &LinkedList.ListNode{}

	less := lessHead
	more := moreHead

	for head != nil {
		if head.Value < x {
			less.Next = head
			less = less.Next
		} else {
			more.Next = head
			more = more.Next
		}
		head = head.Next
	}
	more.Next = nil
	less.Next = moreHead.Next

	return lessHead.Next
}
