package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	seen := make(map[*ListNode]bool)

	curr := head
	seen[curr] = true
	for curr != nil {
		if _, ok := seen[curr.Next]; ok == true {
			return true
		}
		seen[curr] = true
		curr = curr.Next
	}

	return false
}
