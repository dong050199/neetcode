/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
		return 0
	}

	res := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		maxLeft := dfs(node.Left)
		maxRight := dfs(node.Right)
		res = max(res, maxLeft + maxRight)
		return 1 + max(maxLeft, maxRight)
	}
	dfs(root)
	return res
}