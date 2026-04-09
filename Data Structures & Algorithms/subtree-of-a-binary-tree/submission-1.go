func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}

	var dfs func(root, subRoot *TreeNode) bool 
	dfs = func(root, subRoot *TreeNode) bool {
		if root == nil && subRoot == nil {
			return true
		}

		if root == nil || subRoot == nil {
			return false
		}

		if root.Val != subRoot.Val {
			return false
		}

		return dfs(root.Left, subRoot.Left) && dfs(root.Right, subRoot.Right)
	}

	if dfs(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}