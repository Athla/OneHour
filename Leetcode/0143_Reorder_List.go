package leetcode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    middle := func(h *ListNode) *ListNode {
        slow, fast := h, h
        for fast.Next !=nil && fast.Next.Next != nil {
            slow = slow.Next
            fast = fast.Next.Next
        }

        return slow
    }

    reverse := func (h *ListNode) *ListNode {
        var prev, curr, fw *ListNode = nil, h, nil

        for curr != nil {
            fw = curr.Next
            curr.Next = prev
            prev= curr
            curr = fw
        }

        return prev
    } 

    mid := middle(head)
    newHead := mid.Next
    mid.Next = nil

    newHead = reverse(newHead)


    c1 := head
    c2 := newHead
    var f1, f2 *ListNode

    for c1 != nil && c2 != nil {
        f1 = c1.Next
        f2 = c2.Next

        c1.Next = c2
        c2.Next = f1

        c1, c2 = f1, f2
    }
}
