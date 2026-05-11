/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	var dfs func(node *ListNode) bool
	dfs = func(node *ListNode) bool {
		if node == nil {
			return false
		}

		if visited[node] {
			return true
		}
		
		visited[node] = true

		return dfs(node.Next)
	}
	return dfs(head)
}
