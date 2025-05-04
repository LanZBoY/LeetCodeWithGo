package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	dominoCounter := map[string]int{}
	res := 0
	for _, pair := range dominoes {
		cur := fmt.Sprintf("%v_%v", min(pair[0], pair[1]), max(pair[0], pair[1]))
		res += dominoCounter[cur]
		dominoCounter[cur]++
	}
	return res
}

func main() {

	fmt.Printf("%v\n", numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))
	fmt.Printf("%v\n", numEquivDominoPairs([][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}))

}
