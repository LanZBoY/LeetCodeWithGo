package main

import "fmt"

func maxProfit(prices []int) int {
	max_profit := 0
	min_buy := prices[0]

	for i := 1; i < len(prices); i++ {
		max_profit = max(prices[i]-min_buy, max_profit)
		min_buy = min(min_buy, prices[i])
	}
	return max_profit
}

func main() {
	fmt.Printf("%d\n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("%d\n", maxProfit([]int{7, 6, 4, 3, 1}))
}
