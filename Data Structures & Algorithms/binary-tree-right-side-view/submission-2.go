/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    queue := []*TreeNode{root}
	res := []int{}
	newQueue := []*TreeNode{}
	tmp := []int{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		tmp = append(tmp, cur.Val)

		if cur.Left != nil{
			newQueue = append(newQueue, cur.Left)
		}

		if cur.Right != nil {
			newQueue = append(newQueue, cur.Right)
		}

		if len(queue) == 0 {
			queue = newQueue
			res = append(res, tmp[len(tmp)-1])
			newQueue = []*TreeNode{}
			tmp = []int{}
		}
	}

	return res
}
