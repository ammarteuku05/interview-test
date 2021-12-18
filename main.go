package main

import "fmt"

func main() {

	// Max Value
	// n := 5

	// a := []int{0, 1, 2}
	// b := []int{1, 4, 3}

	// k := []int{100, 100, 100}
	// m := len(a)

	// o := FindMax(n, m, a, b, k)

	// fmt.Print(o)

	// Merge two string

	// s1 := "saya"
	// s2 := "kamu"

	// x := MergeTwoString(s1, s2)

	// fmt.Print(x)

	// Morgan And String

	// d := MorganAndString("ABACABA", "ABACABA")

	// fmt.Print(d)

	// JUMP FROG

	// y := JumpFrog(6, 3, 9, 0)
	// fmt.Print(y)

	// Find min max
	// arr := []int{1, 3, 5, 7, 9}

	// x, y := MinMaxSum(arr)

	// fmt.Print(x, y)

	// swap

	// a, b := Swap(3, 5)

	// fmt.Print(a, b)

	// count array appear

	arr := []int{1, 2, 2, 2, 2, 3, 4, 7, 8, 8}
	n := len(arr)
	x := 8

	m := CountAppear(arr, n, x)

	fmt.Print(m)

}
