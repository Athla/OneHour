package dsa

import (
	"log"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
type BST struct {
	Root *TreeNode
}

func (t *BST) Insert(n *TreeNode, v int) *TreeNode {
	if n == nil {
		return &TreeNode{v, nil, nil}
	}

	if v <= t.Root.Val {
		n.Left = t.Insert(n.Left, v)
	}
	if v > n.Val {
		n.Right = t.Insert(n.Right, v)
	}

	return n
}

func (t *BST) Search(n *TreeNode, v int) bool {
	if n.Val == v {
		return true
	}
	if n == nil {
		return false
	}
	if v < n.Val {
		return t.Search(n.Left, v)
	}
	if v > n.Val {
		return t.Search(n.Right, v)
	}
	return false
}

func (t *BST) TraversePreOrder(n *TreeNode) {
	if n == nil {
		return
	} else {
		log.Println(n.Val, " ")
		t.TraversePreOrder(n.Left)
		t.TraversePreOrder(n.Right)
	}
}
func (t *BST) TraverseInOrder(n *TreeNode) {
	if n == nil {
		return
	} else {
		t.TraverseInOrder(n.Left)
		log.Println(n.Val, " ")
		t.TraverseInOrder(n.Right)
	}
}

func (t *BST) TraversePostOrder(n *TreeNode) {
	if n == nil {
		return
	} else {
		t.TraversePostOrder(n.Left)
		t.TraversePostOrder(n.Right)
		log.Println(n.Val, " ")
	}
}
