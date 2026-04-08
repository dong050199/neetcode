/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	cur := root
    if cur == nil {
		return nil
	}

	cur.Left , cur.Right = cur.Right, cur.Left
	invertTree(cur.Left)
	invertTree(cur.Right)
	return root
}
