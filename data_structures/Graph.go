package data_structures

type Graph struct {
	Nodes []int
}

func (g *Graph) AddNode(value int) {
	var node = value
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) AddEdge(n1, n2 GraphNode) {
	n1.Neighbours = append(n1.Neighbours, n2)
	n2.Neighbours = append(n2.Neighbours, n1)
}
