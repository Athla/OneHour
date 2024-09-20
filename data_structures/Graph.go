package data_structures

type GraphNode struct {
	Value      *int
	Neighbours []GraphNode
}

type Graph struct {
	Nodes []GraphNode
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) AddNode(value int) GraphNode {
	var node = GraphNode{Value: &value}
	g.Nodes = append(g.Nodes, node)
	return node
}

func (g *Graph) AddEdge(n1, n2 GraphNode) {
	n1.Neighbours = append(n1.Neighbours, n2)
	n2.Neighbours = append(n2.Neighbours, n1)
}
