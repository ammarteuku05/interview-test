package main

func FindMax(n int, m int, a []int, b []int, k []int) int {
	arr := make([]int, n)

	for i := 0; i < m; i++ {
		lower := a[i]
		upper := b[i]

		for j := lower; j <= upper; j++ {
			arr[j] += k[i]
		}

	}

	var res int
	for i := 0; i < n; i++ {
		res = Max(res, arr[i])
	}

	return res

}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
