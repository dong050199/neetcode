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
		tmp := 0
		if node.Left != nil {
			tmp++
		}

		if node.Right != nil {
			tmp++
		}

		res = max(res, tmp + dfs(node.Left) + dfs(node.Right))

		return tmp
	}

	dfs(root)
	return res
}
