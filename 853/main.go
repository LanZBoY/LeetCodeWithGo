package main

import (
	"fmt"
	"sort"
)

type Car struct {
	position int
	speed    int
}

func carFleet(target int, position []int, speed []int) int {
	size := len(position)
	cars := make([]Car, size)
	for i := range size {
		cars[i].position = position[i]
		cars[i].speed = speed[i]
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	arri_times := make([]float64, 0)

	arri_times = append(arri_times, 0)

	for i := range size {
		arri_time := float64((target - cars[i].position)) / float64(cars[i].speed)
		if arri_times[len(arri_times)-1] < arri_time {
			arri_times = append(arri_times, arri_time)
		}
	}

	return len(arri_times) - 1
}

func main() {
	fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	fmt.Println(carFleet(10, []int{3}, []int{3}))
	fmt.Println(carFleet(100, []int{0, 2, 4}, []int{4, 2, 1}))
}
