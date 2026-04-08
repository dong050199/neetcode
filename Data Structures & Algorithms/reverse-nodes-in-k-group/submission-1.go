func reverseKGroup(head *ListNode, k int) *ListNode {
	var nodes []*ListNode
	cur := head 
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	var revertedNode []*ListNode

	tmp := 0
	var cal []*ListNode
	for _, node := range nodes {
		cal = append(cal, node)
		tmp++
		
		if tmp == k {  
			revertedNode = append(revertedNode, revertNodes(cal)...)
			cal = []*ListNode{}
			tmp = 0
		}
	}

	if len(cal) != 0 {
		revertedNode = append(revertedNode, cal...)  
	}

	dummy := &ListNode{}
	cur = dummy
	
	for _, node := range revertedNode {
		cur.Next = node
		cur = cur.Next
	}
	
	return dummy.Next
}

func revertNodes(nodes []*ListNode) []*ListNode {
	l, r := 0, len(nodes) - 1

	for l <= r {
		nodes[l], nodes[r] = nodes[r], nodes[l]
		l++
		r--
	}

	return nodes
}