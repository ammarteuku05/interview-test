package main

func Swap(a int, b int) (int, int) {
	temp := a
	a = b

	b = temp

	return a, b
}
