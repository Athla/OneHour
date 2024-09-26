package dsa

func (g *Graph) BFS(root int) (order []int) {
	seen := make(map[int]bool)
	q := NewQueue()
	q.Enqueue(root)
	seen[root] = true

	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		order = append(order, v)
		for _, val := range g.Edges[v] {
			if !seen[val] {
				q.Enqueue(val)
				seen[val] = true
			}
		}
	}

	return order
}

// func (g *Graph) BFS(start int) (order []int) {
// 	seen := make(map[int]bool)
// 	q := NewQueue()
// 	q.Enqueue(start)
// 	seen[start] = true
//
// 	for !q.IsEmpty() {
// 		v, _ := q.Dequeue()
// 		order = append(order, v)
// 		for _, val := range g.Edges[v] {
// 			if !seen[val] {
// 				q.Enqueue(val)
// 				seen[val] = true
// 			}
// 		}
// 	}
//
// 	return order
// }
