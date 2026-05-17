/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	var dfs func(node *TreeNode) int 
	res := true
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		maxLeft := dfs(node.Left)
		maxRight := dfs(node.Right)

		if abs(maxLeft - maxRight) > 1 {
			res = false
		}

		return 1 + max(maxLeft, maxRight)
	}

	dfs(root)

	return res
}

func abs(i int) int {
	if i > 0 {
		return i
	}

	return 0 - i
}
