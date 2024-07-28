package main

import (
	"container/heap"
	"fmt"
)

func main() {
	kLarge := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kLarge.Add(3))
}

type KthLargest struct {
	K    int
	Heap *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)

	obj := KthLargest{
		K:    k,
		Heap: h,
	}
	for _, num := range nums {
		obj.Add(num)
	}
	return obj
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.Heap, val)
	if this.Heap.Len() > this.K {
		heap.Pop(this.Heap)
	}
	return (*this.Heap)[0]
}

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
func (h *MinHeap) Push(i any) {
	*h = append(*h, i.(int))
}
func (h *MinHeap) Pop() any {
	old := *h
	length := h.Len()
	x := old[length-1]
	*h = old[0 : length-1]
	return x
}
