package algorithms

import (
	"github.com/Athla/OneHour/data_structures"
	"log"
)

// TODO: FIX GRAPH STRUCT CREATION
func BFS(g data_structures.Graph) []int {
	root := g.Nodes[0]
	seen := make(map[*int]bool)
	order := make([]int, 0)
	q := data_structures.Queue{}

	seen[&root] = true
	q.Enqueue(root)

	for !q.IsEmpty() {
		curr, err := q.Dequeue()
		if err != nil {
			log.Println(err)
		}
		order = append(order, curr)
		if !seen[&curr] {
			seen[&curr] = true
			q.Enqueue(curr)
		}
	}

	return order
}
