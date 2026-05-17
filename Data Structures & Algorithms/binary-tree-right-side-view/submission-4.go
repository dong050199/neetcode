/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

    res := []int{}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		rightNode := nodes[len(nodes)-1]
		res = append(res, rightNode.Val)
		newNodes := []*TreeNode{}
		for _, node := range nodes {
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}

			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}
		nodes = newNodes
	}

	return res
}
