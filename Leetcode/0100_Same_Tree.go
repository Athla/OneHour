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

func isSameTree(p *dsa.TreeNode, q *dsa.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Value != q.Value {
		return false
	}

	l := isSameTree(p.Left, q.Left)
	r := isSameTree(p.Right, q.Right)

	return l && r
}
