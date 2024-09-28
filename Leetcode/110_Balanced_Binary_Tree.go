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
func isBalanced(root *dsa.TreeNode) bool {
	// verify if the curr node is, if it is, return true (because there are no subtres, so diff = 0)
	if root == nil {
		return true
	}

	dif := maxDepth2(root.Left) - maxDepth2(root.Right)

	if dif >= -1 && dif <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func maxDepth2(t *dsa.TreeNode) int {
	if t == nil {
		return 0
	}

	return max(maxDepth(t.Left), maxDepth(t.Right)) + 1
}
