package main

import "fmt"

func largestRectangleArea(heights []int) int {
	size := len(heights)
	maxArea := 0
	for i := range size {
		minHeight := heights[i]
		for j := i; j < size; j++ {
			minHeight = min(minHeight, heights[j])
			maxArea = max(maxArea, minHeight*(j-i+1))
		}
	}
	return maxArea
}

func main() {
	fmt.Printf("%d\n", largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Printf("%d\n", largestRectangleArea([]int{2, 4}))
}
