/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	res := true
	var dfs func(root *TreeNode) int 
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		if abs(left - right) > 1 {
			res = false
		}

		return 1 + max(left, right)
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
