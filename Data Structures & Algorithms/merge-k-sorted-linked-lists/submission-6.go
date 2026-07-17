func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	result := lists[0]
	for i := 1; i < len(lists); i++ {
		result = mergeTwoLinkedList(result, lists[i])
	}
	return result
}

func mergeTwoLinkedList(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy 

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			cur.Next = list1 
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next 
	}

	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}

	return dummy.Next
}