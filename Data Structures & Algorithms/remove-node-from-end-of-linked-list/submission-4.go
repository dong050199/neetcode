/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nodes []*ListNode

	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	if n == len(nodes) {
		if len(nodes) == 1 {
			return nil  // Only node, return nil
		}
		return nodes[1]
	}

	nodes[len(nodes)- 1 - n].Next = nodes[len(nodes) - 1 -n ].Next.Next
	return nodes[0]
}
