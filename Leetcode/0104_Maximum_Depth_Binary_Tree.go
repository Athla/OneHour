package leetcode

import "github.com/Athla/OneHour/dsa"

func maxDepth(root *dsa.TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	if l > r {
		return l + 1
	}
	return r + 1
}
