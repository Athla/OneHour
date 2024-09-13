package data_structures

type Graph struct {
	Nodes []*Node
}

func (g *Graph) AddNode(value int) {
	node := Node{Value: value}
	g.Nodes = append(g.Nodes, &node)
}

func (g *Graph) AddEdge(n1, n2 GraphNode) {
	n1.Neighbours = append(n1.Neighbours, n2)
	n2.Neighbours = append(n2.Neighbours, n1)
}
