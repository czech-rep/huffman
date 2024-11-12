package main

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {

	myheap := &NodeHeap{
		Node{"a", 1, nil, nil},
		Node{"a", 10, nil, nil},
		Node{"a", 11, nil, nil},
		Node{"a", -5, nil, nil},
		Node{"a", 110, nil, nil},
		Node{"a", 1, nil, nil},
		Node{"a", 4, nil, nil},
	}
	heap.Init(myheap)
	heap.Push(myheap, Node{"a", 10, nil, nil})
	heap.Push(myheap, Node{"a", 3, nil, nil})
	heap.Push(myheap, Node{"a", -100, nil, nil})

	if len(*myheap) != 10 {
		t.Errorf("Invalid length of heap")
	}

	minimum := -999 // test that it's a minimum heap
	for len(*myheap) > 0 {
		item := heap.Pop(myheap).(Node)
		if item.Weight < minimum {
			t.Errorf("Heap pop gave %v which is greater than previous %v", item.Weight, minimum)
		}
		minimum = item.Weight
	}
}
