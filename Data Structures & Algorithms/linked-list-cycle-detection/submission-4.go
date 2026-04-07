/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	// using two pointer to detect cycle
	first := head
	second := head

	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next

		if first == second {
			return true
		}
	}

	return false
}
