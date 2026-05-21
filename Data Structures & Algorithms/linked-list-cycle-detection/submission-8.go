/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    visited := make(map[int]bool)
	var dfs func(node *ListNode) 
	res := false
	dfs = func(node *ListNode) {
		if node == nil {
			return
		}

		if visited[node.Val] {
			res = true
			return
		}

		visited[node.Val] = true
		dfs(node.Next)
	}

	dfs(head)

	return res
}
