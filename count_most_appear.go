package main

func CountAppear(arr []int, n int, x int) int {
	res := 0

	for i := 0; i < n; i++ {
		if x == arr[i] {
			res++
		}
	}
	return res
}
