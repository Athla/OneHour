package data_structures

import "errors"

type Queue []int

func (q *Queue) Enqueue(i int) {
	*q = append(*q, i)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue empty.")
	}
	e := (*q)[0]
	*q = (*q)[1:]
	return e, nil
}

func (q *Queue) IsEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}
