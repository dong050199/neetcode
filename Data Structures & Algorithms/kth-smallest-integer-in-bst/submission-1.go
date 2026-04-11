/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	kthSmallest(root.Left, k)
	kthSmallest(root.Right, k)

	res = 1


}
