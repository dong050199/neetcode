func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	nodes := []*ListNode{}

	for head != nil {
		nodes = append(nodes, head)    
		head = head.Next
	}

	i, j := 0, len(nodes)-1

	var tmp bool
	cur := &ListNode{}
	for i <= j {                         
		if !tmp {
			cur.Next = nodes[i]
			i++
		} else {
			cur.Next = nodes[j]
			j--
		}
		cur = cur.Next
		tmp = !tmp                      
	}

	cur.Next = nil                    
	return
}