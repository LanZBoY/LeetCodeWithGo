package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	dominoCounter := map[string]int{}

	for _, pair := range dominoes {
		if pair[0] > pair[1] {
			pair[0], pair[1] = pair[1], pair[0]
		}
		dominoCounter[fmt.Sprintf("%v_%v", pair[0], pair[1])]++
	}

	ret := 0

	for _, v := range dominoCounter {
		ret += (v) * (v - 1) / 2
	}
	return ret
}

func main() {

	fmt.Printf("%v\n", numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))
	fmt.Printf("%v\n", numEquivDominoPairs([][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}))

}
