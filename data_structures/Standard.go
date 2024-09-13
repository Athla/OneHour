package data_structures

func CreateStdGraph() Graph {
	return Graph{Nodes: []*GraphNode{
		{Value: 1, Neighbours: []GraphNode{{Value: 2, Neighbours: []GraphNode{{Value: 3, Neighbours: []GraphNode{{Value: 4}}}}}, {Value: 3, Neighbours: []GraphNode{{Value: 4}}}}},
		{Value: 2, Neighbours: []GraphNode{{Value: 3, Neighbours: []GraphNode{{Value: 4}}}}},
		{Value: 3, Neighbours: []GraphNode{{Value: 4}}},
		{Value: 4},
	}}
}
