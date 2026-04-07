/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	cur := head

	var nodes []*ListNode

	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	l, r := 0, len(nodes) - 1

	for l < r {
		nodes[l].Next = nodes[r]
		l++
		nodes[r].Next = nodes[l]
		r--
	}
	
	nodes[l].Next = nil
}
