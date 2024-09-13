package algorithms

import (
	"github.com/Athla/OneHour/data_structures"
)

func walk(root *data_structures.TreeNode, seen map[*data_structures.TreeNode]bool) (out []int) {
	if root == nil {
		return out
	}
	walk(root.Left)
	return out
}

// Will implement on a tree
func DFS(g *data_structures.Tree) []int {
	seen := make(map[*data_structures.TreeNode]bool)
	s := &data_structures.Stack{}
	if s.IsEmpty() {
		return nil
	}
	root := g.Root
	return walk(root, seen)
}
