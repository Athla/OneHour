package dsa

type Graph struct {
	Edges map[int][]int
}

func NewGraph() *Graph {
	return &Graph{Edges: make(map[int][]int)}
}

func (g *Graph) AddEdge(src, dest int) {
	g.Edges[src] = append(g.Edges[src], dest)
}
