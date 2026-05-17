/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nodes []*ListNode
	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	removeNode := nodes[len(nodes) - n - 1]
	prev := nodes[len(nodes) - n - 2]
	next := removeNode.Next
	prev.Next = next
}
