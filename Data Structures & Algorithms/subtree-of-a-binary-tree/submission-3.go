/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// if isSameTree(root, subRoot) {
	// 	return true
	// } 
	if root == nil {
		return subRoot == nil
	}
	return isSameTree(root.Left, subRoot) || isSameTree(root.Right, subRoot) 
}

func isSameTree(tree1, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 == nil || tree2 == nil {
		return false
	}

	if tree1.Val != tree2.Val {
		return false
	}

	return isSameTree(tree1.Left, tree2.Left) && isSameTree(tree1.Right, tree2.Right)
}
