package linkList

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 != nil {
		return list2
	}
	if list2 == nil && list1 != nil {
		return list1
	}
	if list1 == nil && list2 == nil {
		return nil
	}

	var preHead *ListNode
	pre := &ListNode{}
	preHead = pre
	for list1 != nil && list2 != nil {
		if list1.val <= list2.val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 == nil {
		pre.Next = list2
	}
	if list2 == nil {
		pre.Next = list1
	}
	return preHead.Next
}
