/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Right)
		right := dfs(root.Left)
		res = max(res, left + right)
		return 1 + max(dfs(root.Left), dfs(root.Right))
	}
	dfs(root)
	return res
}
