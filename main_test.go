package main

import (
	"testing"

	"github.com/Athla/OneHour/algorithms"
	"github.com/Athla/OneHour/data_structures"
	"github.com/stretchr/testify/assert"
)

var (
	ExpectedBFS = []int{1, 2, 3, 3, 4, 4}
	Unsorted    = []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9, 15, 18, 22, 22, 19}
	Sorted      = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 18, 19, 22, 22}
)

func TestBFS(T *testing.T) {
	v := data_structures.CreateStdGraph()
	T.Logf("%+v", v)
	order := algorithms.BFS(v)
	T.Logf("Order: %v", order)
	assert.Equal(T, ExpectedBFS, order)
}

func TestQuickSort(t *testing.T) {
	out := algorithms.QuickSort(Unsorted)

	t.Log("\n\tRunning QuickSort")
	if ok := assert.Equal(t, out, Sorted); !ok {
		t.Logf("\nExpected: %v\nReceived %v", Sorted, out)
	} else {
		t.Logf("\nExpected: %v\nReceived %v", Sorted, out)
	}
}

func TestMergeSort(t *testing.T) {
	out := algorithms.MergeSort(Unsorted)

	t.Log("\n\tRunning MergeSort")
	if ok := assert.Equal(t, out, Sorted); !ok {
		t.Logf("\nExpected: %v\nReceived %v", Sorted, out)
	} else {
		t.Logf("\nExpected: %v\nReceived %v", Sorted, out)
	}
}
