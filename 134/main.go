package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	N := len(gas)
	for i := 0; i < N; i++ {
		remainGas := 0
		j := i
		for range N {
			remainGas = remainGas + gas[j]
			remainGas = remainGas - cost[j]

			if remainGas < 0 {
				i = j + 1
				break
			}
			j = (j + 1) % N
		}

		if remainGas >= 0 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Printf("%v\n", canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Printf("%v\n", canCompleteCircuit([]int{6, 0, 1, 3, 2}, []int{4, 5, 2, 5, 5}))
	fmt.Printf("%v\n", canCompleteCircuit([]int{0, 0, 0, 0, 0, 0, 2}, []int{0, 0, 0, 0, 0, 1, 0}))
	fmt.Printf("%v\n", canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	fmt.Printf("%v\n", canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}))
}
