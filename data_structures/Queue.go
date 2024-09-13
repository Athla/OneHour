package data_structures

type Queue []*GraphNode

func (q *Queue) Enqueue(i *GraphNode) {
	*q = append(*q, i)
}

func (q *Queue) Dequeue() *GraphNode {
	if q.IsEmpty() {
		return &GraphNode{}
	}
	e := (*q)[0]
	*q = (*q)[1:]
	return e
}

func (q *Queue) IsEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}
