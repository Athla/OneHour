package data_structures

type Queue []Node

func (q *Queue) Enqueue(i Node) {
	*q = append(*q, i)
}

func (q *Queue) Dequeue() Node {
	if q.IsEmpty() {
		return Node{}
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
