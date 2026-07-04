/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    // create dummy node to use as the result later
	dummy := &ListNode{
		Next: head,
	}

	prev := dummy
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// after this we have the prev node just after left node
	for i := 0; i < right - left; i++ {
		cur := prev.Next
		next = cur.Next
		next 
		prev.Next = next
	}



	return dummy.Next
}
