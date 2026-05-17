/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	cur := head
	oldToNew := make(map[*Node]*Node)

	for cur != nil {
		newNode := &Node{Val: cur.Val}
		oldToNew[cur] = newNode
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		newNode := oldToNew[cur]
		newNode.Next = oldToNew[cur.Next]
		newNode.Random = oldToNew[cur.Random]
		cur = cur.Next
	}

	return oldToNew[head]
}
