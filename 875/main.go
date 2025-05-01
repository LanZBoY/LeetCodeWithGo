package main

import "fmt"

func getTotalHour(piles []int, s int) int {
	sum := 0
	for _, pile := range piles {
		sum += ((pile + s - 1) / s)
	}
	return sum
}

func get_max(piles []int) int {
	maxNum := 0

	for _, pile := range piles {
		maxNum = max(pile, maxNum)
	}

	return maxNum
}

func minEatingSpeed(piles []int, h int) int {
	i := 1
	j := get_max(piles)
	mid := 0
	for i < j {
		mid = (i + j) / 2
		t_h := getTotalHour(piles, mid)
		if t_h <= h {
			j = mid
		} else {
			i = mid + 1
		}
	}

	return i
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}
