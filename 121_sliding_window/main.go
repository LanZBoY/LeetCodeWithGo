package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	l, r := 0, 1
	max_profit := max(0, prices[r]-prices[l])

	for ; r < len(prices); r++ {
		if prices[l] < prices[r] {
			max_profit = max(max_profit, prices[r]-prices[l])
		} else {
			l = r
		}
	}
	return max_profit
}

func main() {
	fmt.Printf("%d\n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("%d\n", maxProfit([]int{7, 6, 4, 3, 1}))
}
