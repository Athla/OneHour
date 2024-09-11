package main

import (
	"testing"

	"github.com/Athla/OneHour/algorithms"
	"github.com/Athla/OneHour/data_structures"
)

func TestBFS(T *testing.T) {
	v := data_structures.CreateStdGraph()
	algorithms.BFS(v)
}
