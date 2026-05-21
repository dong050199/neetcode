/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    visited := make(map[*ListNode]bool)
	var dfs func(node *ListNode) 
	res := false
	dfs = func(node *ListNode) {
		if node == nil {
			return
		}

		if visited[node] {
			res = true
			return
		}

		visited[node] = true
		dfs(node.Next)
	}

	dfs(head)

	return res
}
