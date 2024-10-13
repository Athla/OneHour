package leetcode

import "github.com/Athla/OneHour/dsa"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minDepth(root *dsa.TreeNode) int {
	if root == nil {
		return 0
	}
	low := 1

	q := []*dsa.TreeNode{root}

	for len(q) > 0 {
		le := len(q)

		for i := 0; i < le; i++ {
			curr := q[0]
			q = q[1:]

			if curr.Left == nil && curr.Right == nil {
				return low
			}

			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		low++
	}

	return low
}
