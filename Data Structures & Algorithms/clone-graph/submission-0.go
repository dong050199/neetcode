/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	oldToNew := make(map[*Node]*Node)
	var dfs func (node *Node) *Node 
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		if _, exist := oldToNew[node]; exist {
			return oldToNew[node]
		}

		newNode := &Node{Val: node.Val}
		oldToNew[node] = newNode
		for _, nb := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, dfs(nb))
		}
		return newNode
	}

    return dfs(node)
}
