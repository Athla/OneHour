package algorithms

import (
	"github.com/Athla/OneHour/data_structures"
)

func BFS(g data_structures.Graph) []int {
	root := g.Nodes[0]
	seen := make(map[*data_structures.Node]bool)
	order := make([]int, 0)
	q := data_structures.Queue{}

	seen[&root] = true
	q.Enqueue(root)

	for !q.IsEmpty() {
		curr := q.Dequeue()
		order = append(order, curr.Value)
		for _, v := range curr.Neighbours {
			if !seen[&v] {
				seen[&v] = true
				q.Enqueue(v)
			}
		}
	}

	return order
}
