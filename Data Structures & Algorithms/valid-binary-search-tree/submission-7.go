/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, left, right int) bool
	dfs = func(node *TreeNode, left, right int) bool {
		if node == nil {
			return true
		}

		if node.Val <= left || node.Val >= right {
			return false
		}

		return dfs(node.Left, left, root.Val) && dfs(node.Right, root.Val, right)
	}

	return dfs(root, -1001, 1001)
}
