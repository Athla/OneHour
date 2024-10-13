package leetcode

import "github.com/Athla/OneHour/dsa"

func invertTree(root *dsa.TreeNode) *dsa.TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root

}
