# 02. The Real BFS Please Stand Up, Please Stand Up
- Do you know sometimes when you think you're going REEEEEALLY well at something, just to then... Notice that you're doing the complete opposite of you thought?
- And that's how I, thinking I was implementing BFS, did a horrible version of DFS.
- It happens to the best (or worst) of us, but hey! At least I learned something today.
- Talking about today, it was cool to do this, so far. Took longer than expected but hey, progress is progress.

```go
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
```
# Considerations

- The fun part of it was to implement the helper type of Queue, alongside the Node and Graph.
- I've struggled more to make it work with the defined types than to do the proper implementation, all it took was some reading and fixing.
- To do the check and validation if it's working, I created a function to return a hardcoded graph. This is a part that I think I can improve later, coding something to create those standard inputs, but I'll think about it later.

# Extra code used

```go
type Graph struct {
	Nodes []Node
}

type Node struct {
	Value      int
	Neighbours []Node
}

type Graph struct {
	Nodes []Node
}

// Helper function
func CreateStdGraph() Graph {
	return Graph{Nodes: []Node{
		{Value: 1, Neighbours: []Node{{Value: 2, Neighbours: []Node{{Value: 3, Neighbours: []Node{{Value: 4}}}}}, {Value: 3, Neighbours: []Node{{Value: 4}}}}},
		{Value: 2, Neighbours: []Node{{Value: 3, Neighbours: []Node{{Value: 4}}}}},
		{Value: 3, Neighbours: []Node{{Value: 4}}},
		{Value: 4},
	}}
}
```
