package algorithms

import (
	"github.com/Athla/OneHour/data_structures"
	"log"
)

func Travel(n *data_structures.Node) map[int]*data_structures.Node {
	result := make(map[int]*data_structures.Node)
	if v, ok := result[n.Value]; ok == false {
		result[n.Value] = v
	}
	if n.Neighbours != nil {
		for _, b := range n.Neighbours {
			Travel(b)
		}
	}

	return result
}

func BFS(g *data_structures.Graph) {
	var res map[int]*data_structures.Node
	for _, v := range g.Nodes {
		res = Travel(v)
	}
	log.Println(res)
}
