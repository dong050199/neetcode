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
		return false
	}
	res := true
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, isValid *bool) {
	if root == nil {
		return 
	}

	if root.Left != nil {
		if root.Left.Val >= root.Val {
			*isValid = false
			return
		}
	}

	if root.Right != nil {
		if root.Right.Val <= root.Val {
			*isValid = false
			return
		}
	}

	dfs(root.Right, isValid)	
	dfs(root.Left, isValid)	
}
	