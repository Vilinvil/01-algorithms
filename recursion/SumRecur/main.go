package main

import (
	"fmt"
	"time"
)

// SumRecur calculate sum of all elements of []int. Func use recursion
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

// SumIter  calculate sum of all elements of []int. Func use cycles
func SumIter(sl []int) int {
	var sum int
	for _, val := range sl {
		sum += val
	}
	return sum
}

func main() {
	a := make([]int, 0)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}

	t0 := time.Now()
	fmt.Println(SumRecur(a))
	t1 := time.Now()
	fmt.Println(t1.Sub(t0))

	t0 = time.Now()
	fmt.Println(SumIter(a))
	t1 = time.Now()
	fmt.Println(t1.Sub(t0))
}
