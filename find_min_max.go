package main

import (
	"sort"
)

func MinMaxSum(arr []int) (int, int) {

	arr2 := make([]int, 0)

	for _, v := range arr {
		arr2 = append(arr2, int(v))
	}
	sort.Ints(arr2)
	var min, max int

	for i, v := range arr2 {
		if i > 3 {
			break
		}
		min += v
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr2)))
	for i, v := range arr2 {
		if i > 3 {
			break
		}
		max += v
	}
	return min, max

}
