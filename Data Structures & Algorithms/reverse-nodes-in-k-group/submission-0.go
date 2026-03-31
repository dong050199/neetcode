/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    cur := head
	group := 0

	for cur != nil && group < k {
		cur = cur.Next
		group ++
	}

	if group == k {
		cur = reverseKGroup(cur, k)
		for group > 0 {
			tmp := head.Next
			head.Next = cur
			cur = head
			head = tmp
			group--
		}
		head = cur
	}

	return head
}
