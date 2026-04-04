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

	var tmp []*ListNode{}

	// remove the n-th node from array
	tmp = append(tmp, nodes[:len(nodes) - n - 1]...)
	tmp = append(tmp, nodes[len(nodes) - n]...)

	// re-create linked list from array
	

	
	
}
