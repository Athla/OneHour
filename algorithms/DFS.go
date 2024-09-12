package algorithms

import (
	"github.com/Athla/OneHour/data_structures"
)

func Travel(n *data_structures.Node) (map[*data_structures.Node]int, []int) {
	result := make(map[*data_structures.Node]int)
	var order []int
	if v, ok := result[n]; !ok {
		result[n] = v
		order = append(order, n.Value)
	}
	if n.Neighbours != nil {
		for _, b := range n.Neighbours {
			Travel(&b)
		}
	}

	return result, order
}

func DFS(g *data_structures.Graph) {
}
