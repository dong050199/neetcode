/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	var dfs func(root *TreeNode, left, right int) bool
	dfs = func(root *TreeNode, left, right int) bool {
		if root == nil {
			return true
		}

		if root.Val <= left || root.Val >= right {
			return false
		}

		return dfs(root.Left, left, root.Val) && dfs(root.Right, root.Val, right)
	}


	return dfs(root, -10000, 10000)
}
