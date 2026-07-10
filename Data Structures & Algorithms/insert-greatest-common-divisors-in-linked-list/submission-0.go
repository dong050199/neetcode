/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    dummy := &ListNode{
        Next: head,
    }
    
    cur := dummy.Next
    for cur != nil && cur.Next != nil {
        next := cur.Next
        gcd := GCD(cur.Val, next.Val)
        newNode := &ListNode{
            Val: gcd,
            Next: next,
        }

        cur.Next = newNode
        cur = next
    }
    return dummy.Next
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}