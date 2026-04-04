/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    nodes := []*ListNode{}

	cur := head
	for cur != nil {
		nodes = append(nodes, curr)
		cur = cur.Next
	}

	removeIndex := len(nodes) - n
	// return immediate if the remove index is 0
	if removeIndex == 0 {
		return head.Next
	}

	// update next of just before remove index to the next of removed node
	nodes[removeIndex-1].Next = nodes[removeIndex].Next
	return head
}
