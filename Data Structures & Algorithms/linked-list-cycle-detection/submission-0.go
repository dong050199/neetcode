/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	mp := make(map[*ListNode]bool)

	for head != nil {
		if mp[head] {
			return true
		}
		mp[head] = true
		head = head.Next
	}

	return false
}
