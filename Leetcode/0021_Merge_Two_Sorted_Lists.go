package leetcode

// Goes until there is only one node and returns
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Stop condition from recursion
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// Verify the values. Call the next value of the smaller one, passing the unchanged list alongside
	if list2.Val > list1.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	// Which means list1.Val > list2.Val
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
