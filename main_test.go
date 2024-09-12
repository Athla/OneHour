package main

import (
	"slices"
	"testing"

	"github.com/Athla/OneHour/algorithms"
	"github.com/Athla/OneHour/data_structures"
)

var (
	ExpectedBFS = []int{1, 2, 3, 3, 4, 4}
)

func TestBFS(T *testing.T) {
	v := data_structures.CreateStdGraph()
	T.Logf("%+v", v)
	order := algorithms.BFS(v)
	T.Logf("Order: %v", order)
	if !slices.Equal(order, ExpectedBFS) {
		T.Errorf("Expected: %v \tReceived: %v", ExpectedBFS, order)
	}
}
