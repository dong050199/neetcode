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
	tmp := []int{}
	queue := []*TreeNode{root}
	newQueue := []*TreeNode{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		tmp = append(tmp, cur.Val)
		if cur.Left != nil {
			newQueue = append(newQueue, cur.Left)
		}
		if cur.Right != nil {
			newQueue = append(newQueue, cur.Right)
		}
		if len(queue) == 0 {
			res = append(res, tmp)
			queue = newQueue
			newQueue = []*TreeNode{}
			tmp = []int{}
		}
	}

	return res
}
