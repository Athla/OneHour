package data_structures

type GraphNode struct {
	Value      *Node
	Neighbours []GraphNode
}

type TreeNode struct {
	Val, Right, Left *Node
}

type Node struct {
	Value int
}
