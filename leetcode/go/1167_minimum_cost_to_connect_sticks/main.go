package main

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func connectSticks(sticks []int) int {
	h := &IntHeap{}

	heap.Init(h)

	for _, stick := range sticks {
		heap.Push(h, stick)
	}

	sum := 0
	for h.Len() > 1 {
		first, second := heap.Pop(h).(int), heap.Pop(h).(int)

		sum += first + second
		heap.Push(h, first+second)
	}

	return sum
}
