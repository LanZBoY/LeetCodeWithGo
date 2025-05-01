package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m := len(nums1)
	n := len(nums2)
	partial_point := (m + n + 1) / 2

	left := 0
	right := m
	left_max := 0
	right_min := 0
	for left <= right {
		i := (left + right) / 2
		j := partial_point - i
		// Process Left Part
		left_1 := nums1[:i]
		left_num_1 := math.MinInt32
		if len(left_1) > 0 {
			left_num_1 = left_1[len(left_1)-1]
		}

		left_2 := nums2[:j]
		left_num_2 := math.MinInt32
		if len(left_2) > 0 {
			left_num_2 = left_2[len(left_2)-1]
		}

		left_max = max(left_num_1, left_num_2)
		// Process Right Part
		right_1 := nums1[i:]
		right_num_1 := math.MaxInt32
		if len(right_1) > 0 {
			right_num_1 = right_1[0]
		}

		right_2 := nums2[j:]
		right_num_2 := math.MaxInt32
		if len(right_2) > 0 {
			right_num_2 = right_2[0]
		}

		right_min = min(right_num_1, right_num_2)

		if left_num_1 > right_num_2 {
			right = i - 1
		} else if left_num_2 > right_num_1 {
			left = i + 1
		} else {
			break
		}
	}
	if (m+n)%2 == 0 {
		return (float64(left_max) + float64(right_min)) / 2
	}

	return float64(left_max)
}

func main() {
	fmt.Printf("%.2f\n", findMedianSortedArrays([]int{4, 5}, []int{1, 2, 3}))
	fmt.Printf("%.2f\n", findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Printf("%.2f\n", findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Printf("%.2f\n", findMedianSortedArrays([]int{5, 7, 8, 11}, []int{4, 9, 12, 14}))
	fmt.Printf("%.2f\n", findMedianSortedArrays([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}))
	fmt.Printf("%.2f\n", findMedianSortedArrays([]int{1, 3}, []int{2, 4, 5, 6, 7, 8}))
	fmt.Printf("%.2f\n", findMedianSortedArrays([]int{2, 2, 4, 4}, []int{2, 2, 2, 4, 4}))
}
