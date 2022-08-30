package main

import "fmt"

// FindMaxNum find max element in []int and return its value
func FindMaxNum(sl []int, max int) int {
	switch len(sl) {
	case 1:
		{
			if sl[0] >= max {
				return sl[0]
			} else {
				return max
			}
		}
	default:
		{
			if sl[0] >= max {
				return FindMaxNum(sl[1:], sl[0])
			} else {
				return FindMaxNum(sl[1:], max)
			}
		}
	}
	//return max
}

func main() {
	a := []int{-1, 2, -3, -4, -5, -6, 7, 18, -9}
	fmt.Println(FindMaxNum(a, a[0]))
}
