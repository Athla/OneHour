package dsa

func CreateStdGraph() *Graph {
	g := &Graph{
		Edges: map[int][]int{
			0: {1, 2},
			1: {3, 4},
			2: {},
			3: {},
			4: {},
		},
	}

	return g
}
