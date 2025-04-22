package main

import "fmt"

func prevSmallerElements(heights []int) []int {
	maxStack := make([]int, 0)
	prevSmallerElements := make([]int, len(heights))

	for i := range len(heights) {
		for len(maxStack) > 0 && heights[maxStack[len(maxStack)-1]] >= heights[i] {
			maxStack = maxStack[:len(maxStack)-1]
		}
		if len(maxStack) > 0 {
			prevSmallerElements[i] = maxStack[len(maxStack)-1]
		} else {
			prevSmallerElements[i] = -1
		}
		maxStack = append(maxStack, i)
	}
	return prevSmallerElements
}

func nextSmallerElements(heights []int) []int {
	maxStack := make([]int, 0)
	nextSmallerElements := make([]int, len(heights))

	for i := len(nextSmallerElements) - 1; i >= 0; i-- {
		for len(maxStack) > 0 && heights[maxStack[len(maxStack)-1]] >= heights[i] {
			maxStack = maxStack[:len(maxStack)-1]
		}
		if len(maxStack) > 0 {
			nextSmallerElements[i] = maxStack[len(maxStack)-1]
		} else {
			nextSmallerElements[i] = len(heights)
		}
		maxStack = append(maxStack, i)
	}
	return nextSmallerElements
}

func largestRectangleArea(heights []int) int {
	pse := prevSmallerElements(heights)
	nse := nextSmallerElements(heights)
	maxArea := 0
	for i := range len(heights) {
		maxArea = max(maxArea, (nse[i]-pse[i]-1)*heights[i])
	}
	return maxArea
}

func main() {
	fmt.Printf("%d\n", largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Printf("%d\n", largestRectangleArea([]int{2, 4}))
}
