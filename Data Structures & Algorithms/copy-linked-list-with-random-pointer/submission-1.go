func copyRandomList(head *Node) *Node {
	oldToNew := make(map[*Node]*Node)
	
	cur := head
	for cur != nil {
		newNode := &Node{Val: cur.Val}
		oldToNew[cur] = newNode
		cur = cur.Next
	}

	// Second pass: fix pointers using the map
	cur = head
	for cur != nil {
		newNode := oldToNew[cur]
		newNode.Next = oldToNew[cur.Next]   // map lookup returns nil if key doesn't exist
		newNode.Random = oldToNew[cur.Random]
		cur = cur.Next
	}

	return oldToNew[head]
}