package algorithms

import (
	"github.com/Athla/OneHour/data_structures"
)

func walk(root *int, seen map[*data_structures.TreeNode]bool) (out []int) {
	if root == nil {
		return out
	}
	walk(root, seen)
	return out
}

// Will implement on a tree
func DFS(g *data_structures.BST) []int {
	seen := make(map[*data_structures.TreeNode]bool)
	s := &data_structures.Stack{}
	if s.IsEmpty() {
		return nil
	}
	root := g.Root
	return walk(&root.Value, seen)
}
