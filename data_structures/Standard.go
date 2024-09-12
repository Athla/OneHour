package data_structures

func CreateStdGraph() Graph {
	return Graph{Nodes: []Node{
		{Value: 1, Neighbours: []Node{{Value: 2, Neighbours: []Node{{Value: 3, Neighbours: []Node{{Value: 4}}}}}, {Value: 3, Neighbours: []Node{{Value: 4}}}}},
		{Value: 2, Neighbours: []Node{{Value: 3, Neighbours: []Node{{Value: 4}}}}},
		{Value: 3, Neighbours: []Node{{Value: 4}}},
		{Value: 4},
	}}
}
