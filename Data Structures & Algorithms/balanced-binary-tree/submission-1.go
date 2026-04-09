/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	res := true

	var dfs func(root *TreeNode) int 
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		if abs(left, right) > 1 {
			res = false
			return 0
		}

		return 1 + max(left, right)
	}

	dfs(root)

	return res
}

func abs(a , b int) int {
	if a > b {
		return a - b
	} else if a < b {
		return b - a
	} else {
		return 0
	}
}
