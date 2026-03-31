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

	isBalance := true
	dfs(root, &isBalance)

	return isBalance
}

func dfs(root *TreeNode,isBalance *bool) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left, isBalance)
	right := dfs(root.Right, isBalance)

	diff := left - right

	if diff > 1 || diff < -1 {
		*isBalance = false
	}

	return 1 + max(left, right)
}
