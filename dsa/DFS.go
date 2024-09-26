package dsa

func (g *Graph) DFSWalk(root int, seen map[int]bool, order *[]int) {
	seen[root] = true
	*order = append(*order, root)
	for _, v := range g.Edges[root] {
		if !seen[v] {
			g.DFSWalk(v, seen, order)
		}
	}
}

func (g *Graph) DFS(root int) []int {
	seen := make(map[int]bool)
	order := []int{}

	g.DFSWalk(root, seen, &order)
	return order
}
