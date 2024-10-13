package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	l := mergeSort(arr[len(arr)/2:])
	r := mergeSort(arr[:len(arr)/2])

	return merge(l, r)
}

func merge(a, b []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		result = append(result, a[i])
	}

	for ; j < len(b); j++ {
		result = append(result, b[j])
	}

	return result
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	val := []int{}
	curr := head
	for curr != nil {
		val = append(val, curr.Val)
		curr = curr.Next
	}

	sortVal := mergeSort(val)
	var resNode *ListNode = &ListNode{Val: sortVal[0], Next: nil}
	root := resNode

	for i := 1; i < len(sortVal); i++ {
		resNode.Next = &ListNode{Val: sortVal[i], Next: nil}
		resNode = resNode.Next
	}

	return root

}
