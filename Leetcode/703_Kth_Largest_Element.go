package leetcode

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type KthLargest struct {
	elem *MinHeap
	size int
}

func constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)
	large := KthLargest{
		elem: h,
		size: k,
	}
	for _, v := range nums {
		large.Add(v)
	}
	return large
}

func (k *KthLargest) Add(val int) int {
	if k.elem.Len() < k.size {
		heap.Push(k.elem, val)
	} else if val > (*k.elem)[0] {
		heap.Pop(k.elem)
		heap.Push(k.elem, val)
	}

	if k.elem.Len() > 0 {
		return (*k.elem)[0]
	}
	return 0
}
