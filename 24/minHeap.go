package main

import (
	"container/heap"

	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] []T

func (h MinHeap[T]) Len() int           { return len(h) }
func (h MinHeap[T]) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap[T]) Push(x any)        { *h = append(*h, x.(T)) }

func (h *MinHeap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap[T]) Top() T {
	if len(h) == 0 {
		panic("heap is empty")
	}
	return h[0]
}

func GetMinHeap[T constraints.Ordered]() *MinHeap[T] {
	h := &MinHeap[T]{}
	heap.Init(h)
	return h
}
