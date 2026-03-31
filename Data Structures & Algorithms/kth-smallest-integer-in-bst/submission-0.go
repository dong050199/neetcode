/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    // travel order reversal 
	arr := []int{}
	dfs(root, &arr)
	return arr[k-1]


}

func dfs(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, arr)
	*arr = append(*arr, root.Val)
	dfs(root.Right, arr)
}


