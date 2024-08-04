package LinkedList

// ListNode represents a node in a linked list.
type ListNode struct {
	Value int
	Next  *ListNode
}

// New creates a new ListNode with the given value.
func New(d int) *ListNode {
	n := new(ListNode)
	n.Value = d
	return n
}

// Length returns the length of the linked list
func (head *ListNode) Length() int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}

// CreateList converts a slice of integers into a linked list.
func CreateList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := New(values[0])
	current := head
	for _, val := range values[1:] {
		current.Next = New(val)
		current = current.Next
	}
	return head
}

// ListsEqual compares two linked lists for equality.
func ListsEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Value != l2.Value {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
