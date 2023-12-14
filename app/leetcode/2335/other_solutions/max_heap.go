package other_solutions

import "container/heap"

func FillCupsByMaxHeap(amount []int) int {
	result := 0

	// initialize the heap data structure
	h := &IntHeap{}

	// add all the values to heap, O(n log n)
	for _, val := range amount { // O(n)
		heap.Push(h, val) // O(log n)
	}

	var max, secondMax int
	for h.Len() > 0 {
		max = heap.Pop(h).(int)
		if max > 0 {
			result++
		}
		if h.Len() > 0 {
			secondMax = heap.Pop(h).(int)
		}

		if max-1 > 0 {
			heap.Push(h, max-1)
		}
		if secondMax-1 > 0 {
			heap.Push(h, secondMax-1)
		}
	}

	return result
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
