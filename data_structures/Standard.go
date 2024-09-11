package data_structures

func CreateStdGraph() *Graph {
	graph := &Graph{}

	// Create nodes
	node1 := graph.AddNode(1)
	node2 := graph.AddNode(2)
	node3 := graph.AddNode(3)
	node4 := graph.AddNode(4)
	node5 := graph.AddNode(5)

	graph.AddEdge(node1, node2)
	graph.AddEdge(node1, node3)
	graph.AddEdge(node2, node4)
	graph.AddEdge(node3, node4)
	graph.AddEdge(node3, node5)
	graph.AddEdge(node4, node5)

	return graph
}
