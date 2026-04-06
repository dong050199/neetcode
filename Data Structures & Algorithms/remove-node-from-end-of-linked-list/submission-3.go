func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{}
	dummy.Next = head
	left := dummy
	right := dummy
	
	// Move right pointer n+1 steps ahead
	for i := 0; i <= n; i++ {
		right = right.Next
	}

	// Move both pointers until right is nil
	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

	return dummy.Next
}