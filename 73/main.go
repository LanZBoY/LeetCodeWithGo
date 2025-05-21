package main

func setZeroes(matrix [][]int) {
	addressSet := map[*int]bool{}
	M := len(matrix)
	N := len(matrix[0])

	for i := range M {
		for j := range N {
			if matrix[i][j] != 0 {
				continue
			}
			addressSet[&matrix[i][j]] = true

			for temp := range M {
				addressSet[&matrix[temp][j]] = true
			}

			for temp := range N {
				addressSet[&matrix[i][temp]] = true
			}
		}
	}

	for k := range addressSet {
		*k = 0
	}
}

func main() {
	setZeroes([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	})
	setZeroes([][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	})
}
