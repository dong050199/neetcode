/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	maxDepth := 0
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)
		maxDepth = max(maxDepth, left + right)

		return 1 + max(left, right)
	}

	dfs(root)

	return maxDepth
}
