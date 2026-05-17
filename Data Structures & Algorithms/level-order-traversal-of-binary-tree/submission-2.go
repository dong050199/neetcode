/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
    nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		newNodes := []*TreeNode{}
		arr := []int{}
		for _, node := range nodes {
			arr = append(arr, node.Val)
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}

			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}

		nodes = newNodes
		res = append(res, arr)
	}

	return res
}
