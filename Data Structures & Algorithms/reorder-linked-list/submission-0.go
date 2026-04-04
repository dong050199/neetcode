/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    var nodes []*ListNode

	curr = head
	for curr != nil {
		nodes = append(nodes, curr)
		curr.Next
	}

	l, r := 0, len(nodes) - 1



}
