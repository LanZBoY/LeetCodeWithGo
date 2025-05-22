package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func maxRemoval(nums []int, queries [][]int) int {
	// 根據第一個來sort
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})

	maxHeap := &Heap{}
	heap.Init(maxHeap)
	deltaArr := make([]int, len(nums)+1)
	operations := 0

	for i, j := 0, 0; i < len(nums); i++ {
		operations += deltaArr[i]
		for j < len(queries) && queries[j][0] == i {
			heap.Push(maxHeap, queries[j][1])
			j++
		}

		for operations < nums[i] && maxHeap.Len() > 0 && (*maxHeap)[0] >= i {
			operations++
			deltaArr[heap.Pop(maxHeap).(int)+1]--
		}

		if operations < nums[i] {
			return -1
		}
	}
	return maxHeap.Len()
}

func main() {
	fmt.Printf("%v\n", maxRemoval([]int{2, 0, 2}, [][]int{{0, 2}, {0, 2}, {1, 1}}))
	fmt.Printf("%v\n", maxRemoval([]int{1, 1, 1, 1}, [][]int{{1, 3}, {0, 2}, {1, 3}, {1, 2}}))
	fmt.Printf("%v\n", maxRemoval([]int{1, 2, 3, 4}, [][]int{{0, 3}}))
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i int, j int) bool {
	return h[i] > h[j]
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h Heap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}
