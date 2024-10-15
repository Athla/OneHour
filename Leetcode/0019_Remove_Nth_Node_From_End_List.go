package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{0, head}
	lead := res

	for i := 0; i <= n; i++ {
		lead = lead.Next
	}

	curr := res
	for lead != nil {
		curr = curr.Next
		lead = lead.Next
	}
	curr.Next = curr.Next.Next

	return res.Next
}
