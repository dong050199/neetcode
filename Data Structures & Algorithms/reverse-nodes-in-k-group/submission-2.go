func reverseKGroup(head *ListNode, k int) *ListNode {
	var nodes []*ListNode
	cur := head 
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	var revertedNode []*ListNode

	for i := 0; i < len(nodes); i += k {
		end := i + k
		if end > len(nodes) {
			end = len(nodes)
		}
		
		group := nodes[i:end]
		if len(group) == k {
			revertedNode = append(revertedNode, revertNodes(group)...)
		} else {
			revertedNode = append(revertedNode, group...)
		}
	}

	dummy := &ListNode{}
	cur = dummy
	
	for _, node := range revertedNode {
		cur.Next = node
		cur = node
	}
	cur.Next = nil  // ← Clear the last node's Next
	
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