/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {

	cur := root

	for cur != nil {
		switch {
			case cur.Val > p.Val && cur.Val > q.Val:
				cur = cur.Left
			case  cur.Val < p.Val && cur.Val < q.Val:
				cur = cur.Right
			default:
				return cur
		}
	}
	return nil
}
