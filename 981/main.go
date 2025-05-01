package main

import "fmt"

type TimeMap struct {
	data          map[string]map[int]string
	key2TimeOrder map[string][]int
}

func Constructor() TimeMap {
	return TimeMap{
		data:          map[string]map[int]string{},
		key2TimeOrder: map[string][]int{},
	}
}

func (obj *TimeMap) Set(key string, value string, timestamp int) {
	if obj.data[key] == nil {
		obj.data[key] = map[int]string{}
		obj.key2TimeOrder[key] = []int{}
	}
	obj.data[key][timestamp] = value
	obj.key2TimeOrder[key] = append(obj.key2TimeOrder[key], timestamp)
}

func binarySearch(timeOrder []int, target int) int {
	i := 0
	j := len(timeOrder)

	if target < timeOrder[i] {
		return -1
	}

	if target > timeOrder[len(timeOrder)-1] {
		return len(timeOrder) - 1
	}

	mid := 0
	for i < j {
		mid = (i + j) / 2

		if mid > target {
			j = mid - 1
		} else if mid < target {
			i = mid + 1
		} else {
			return mid
		}
	}

	return max(0, mid-1)
}

func (obj *TimeMap) Get(key string, timestamp int) string {
	keyMap := obj.data[key]
	if keyMap == nil {
		return ""
	}

	if val, ok := keyMap[timestamp]; ok {
		return val
	}

	timeOrder := obj.key2TimeOrder[key]
	index := binarySearch(timeOrder, timestamp)

	if index == -1 {
		return ""
	}

	return keyMap[timeOrder[index]]
}

func main() {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)
	timeMap.Set("foo", "bar4", 4)
	fmt.Println(timeMap.Get("foo", 3))
	fmt.Println(timeMap.Get("foo", 5))
	// timeMap.Set("love", "high", 10)
	// timeMap.Set("love", "low", 20)
	// fmt.Println(timeMap.Get("love", 5))
	// fmt.Println(timeMap.Get("love", 10))
	// fmt.Println(timeMap.Get("love", 15))
	// fmt.Println(timeMap.Get("love", 20))
}
