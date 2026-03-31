func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    queue := []*TreeNode{root}
    var res [][]int

    for len(queue) > 0 {
        k := len(queue) 
        var l []int
        
        for i := 0; i < k; i++ {
            cur := queue[0]
            queue = queue[1:] 

            l = append(l, cur.Val)

            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
        res = append(res, l)
    }

    return res
}