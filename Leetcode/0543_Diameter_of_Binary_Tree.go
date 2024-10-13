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
func diameterOfBinaryTree(root *dsa.TreeNode) int {
	var diameter int
	if root == nil {
		return 0
	}
	var walk func(n *dsa.TreeNode) int
	walk = func(n *dsa.TreeNode) int {
		if n == nil {
			return 0
		}
		l := walk(n.Left)
		r := walk(n.Right)

		diameter = max(diameter, (l + r))

		return max(l, r) + 1
	}

	walk(root)

	return diameter
}

// func max(a,  b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }
