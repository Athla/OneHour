package leetcode

import "github.com/Athla/OneHour/dsa"

func isSymmetric(root *dsa.TreeNode) bool {
	if root == nil {
		return true
	}

	qL := []*dsa.TreeNode{root.Left}
	qR := []*dsa.TreeNode{root.Right}

	for len(qL) > 0 && len(qR) > 0 {
		l := qL[0]
		r := qR[0]

		qL = qL[1:]
		qR = qR[1:]

		if l == nil && r == nil {
			continue
		}

		if l == nil || r == nil || l.Val != r.Val {
			return false
		}

		qL = append(qL, l.Left, l.Right)
		qR = append(qR, r.Right, r.Left)

	}

	return true
}
