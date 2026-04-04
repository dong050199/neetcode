/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr = head 
	var prev *ListNode // empty node as the the previous node of the current head
	// iterate until we reach the end of linked list
	for curr != nil {
		// store the current next
		tmp := curr.Next
	}
	return prev
}