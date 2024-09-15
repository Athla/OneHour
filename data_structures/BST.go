package data_structures

import (
	"log"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}
type BST struct {
	Root *TreeNode
}

func (t *BST) Insert(n *TreeNode, v int) *TreeNode {
	if t.Root == nil {
		t.Root = &TreeNode{v, nil, nil}
	}
	if n == nil {
		return &TreeNode{v, nil, nil}
	}

	if v <= t.Root.Value {
		n.Left = t.Insert(n.Left, v)
	}
	if v > n.Value {
		n.Left = t.Insert(n.Right, v)
	}

	return n
}

func (t *BST) Search(n *TreeNode, v int) bool {
	if n.Value == v {
		return true
	}
	if n == nil {
		return false
	}
	if v < n.Value {
		return t.Search(n.Left, v)
	}
	if v > n.Value {
		return t.Search(n.Right, v)
	}
	return false
}

func (t *BST) TraversePreOrder(n *TreeNode) {
	if n == nil {
		return
	} else {
		log.Println(n.Value, " ")
		t.TraverseInOrder(n.Left)
		t.TraverseInOrder(n.Right)
	}
}
func (t *BST) TraverseInOrder(n *TreeNode) {
	if n == nil {
		return
	} else {
		t.TraverseInOrder(n.Left)
		log.Println(n.Value, " ")
		t.TraverseInOrder(n.Right)
	}
}

func (t *BST) TraversePostOrder(n *TreeNode) {
	if n == nil {
		return
	} else {
		t.TraverseInOrder(n.Left)
		t.TraverseInOrder(n.Right)
		log.Println(n.Value, " ")
	}
}
