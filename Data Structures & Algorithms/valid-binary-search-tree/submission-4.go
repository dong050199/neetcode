/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && root.Right != nil {
		// validate leaf
		if root.Left.Val >= root.Right.Val {
			return false
		}
		// validate root and leaf
		if root.Val <= root.Left.Val || root.Val >= root.Right.Val {
			return false
		}
	}

	if root.Left != nil {
		return isValidBST(root.Left)
	}

	if root.Right != nil {
		return isValidBST(root.Right)
	}

	return  isValidBST(root.Left) && isValidBST(root.Right)
}
