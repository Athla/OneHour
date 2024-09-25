package dsa

func (g *Graph) DFS(root int) []int {
	seen := make(map[int]bool)
	order := []int{}
	g.Walk(root, seen, &order)
	return order
}

func (g *Graph) Walk(root int, seen map[int]bool, order *[]int) {
	seen[root] = true
	*order = append(*order, root)

	for _, v := range g.Edges[root] {
		if !seen[v] {
			g.Walk(v, seen, order)
		}
	}
}

// func (g *Graph) DFSUtil(s int, seen map[int]bool, order *[]int) {
// 	seen[s] = true
// 	*order = append(*order, s)
// 	for _, v := range g.Edges[s] {
// 		if !seen[v] {
// 			g.DFSUtil(v, seen, order)
// 		}
// 	}
// }
//
// func (g *Graph) DFS(start int) []int {
// 	seen := make(map[int]bool)
// 	order := make([]int, 0)
// 	g.DFSUtil(start, seen, &order)
// 	return order
// }
