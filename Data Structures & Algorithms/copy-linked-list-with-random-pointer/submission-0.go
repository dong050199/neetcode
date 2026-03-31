/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	// create copy of all nodes from linked list
	mp := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		newNode := &Node{
			Val: cur.Val,
		}
		mp[cur] = newNode
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		newNode := mp[cur]
		newNode.Next = mp[cur.Next]
		newNode.Random = mp[cur.Random]
		cur = cur.Next
	}

	return mp[head]
}
