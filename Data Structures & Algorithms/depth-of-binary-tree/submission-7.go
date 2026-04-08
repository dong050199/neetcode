/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	cur := root
	if cur == nil {
		return 0
	}

	return 1 + max(maxDepth(cur.Left), maxDepth(cur.Right))
}
