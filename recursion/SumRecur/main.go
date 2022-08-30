package main

import "fmt"

func SumRecur(sl []int) int {
	switch len(sl) {
	case 0:
		return 0

	case 1:
		return sl[0]

	default:
		return sl[0] + SumRecur(sl[1:])
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(SumRecur(a))
}
