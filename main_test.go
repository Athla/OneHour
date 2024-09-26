package main

import (
	"container/heap"
	"testing"

	data_structures "github.com/Athla/OneHour/dsa"
	"github.com/stretchr/testify/assert"
)

var (
	ExpectedBFS   = []int{1, 2, 3, 4, 7, 5, 6}
	ExpectedDFS   = []int{1, 2, 7, 3, 4, 5, 6}
	Unsorted      = []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9, 15, 18, 22, 22, 19}
	Sorted        = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 18, 19, 22, 22}
	StandardGraph = &data_structures.Graph{
		Edges: map[int][]int{
			1: {2, 3, 4},
			2: {7},
			3: {4, 7},
			4: {5, 6},
			5: {},
			6: {},
			7: {},
		},
	}
	ExpectedHeap = &data_structures.MinHeap{1, 2, 5, 3}
)

func TestBinarySearch(T *testing.T) {
	arr := []int{0, 1, 2, 3, 5, 6, 7, 8, 9, 10}

	ok, idx := data_structures.BinarySearch(arr, 10)
	assert.Equal(T, ok, true)
	assert.Equal(T, idx, 9)

	ok, idx = data_structures.BinarySearch(arr, 4)
	assert.Equal(T, ok, false)
	assert.Equal(T, idx, -1)
}
func TestBFS(T *testing.T) {
	bfs := StandardGraph.BFS(1)
	T.Log("\nBFS traversal starting from vertex 0:", bfs)
	assert.Equal(T, bfs, ExpectedBFS)
}

func TestDFS(T *testing.T) {
	dfs := StandardGraph.DFS(1)
	T.Log("\nDFS traversal starting from vertex 0:", dfs)
	assert.Equal(T, dfs, ExpectedDFS)
}

func TestQuickSort(t *testing.T) {
	out := data_structures.QuickSort(Unsorted)

	t.Log("\n\tRunning QuickSort")
	if ok := assert.Equal(t, out, Sorted); !ok {
		t.Logf("\nExpected: %v\nReceived %v", Sorted, out)
	} else {
		t.Logf("\nExpected: %v\nReceived %v", Sorted, out)
	}
}

func TestMergeSort(t *testing.T) {
	out := data_structures.MergeSort(Unsorted)

	t.Log("\n\tRunning MergeSort")
	if ok := assert.Equal(t, out, Sorted); !ok {
		t.Logf("\nExpected: %v\nReceived %v", Sorted, out)
	} else {
		t.Logf("\nExpected: %v\nReceived %v", Sorted, out)
	}
}

func TestHeap(t *testing.T) {
	h := &data_structures.MinHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	t.Logf("Current heap: %v ----------- Expectec head: %v", h, ExpectedHeap)
	assert.Equal(t, h, ExpectedHeap)
}
