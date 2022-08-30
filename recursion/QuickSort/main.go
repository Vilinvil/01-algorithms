package main

import (
	"fmt"
)

// QuickSort ascending(по возрастанию)
func QuickSort(sl []int) []int {
	if len(sl) <= 1 {
		return sl
	} else {
		mid := (len(sl) - 1) / 2

		var left, right []int
		for i := 0; i < len(sl); i++ {
			if i == mid {
				i++
			}
			if sl[i] < sl[mid] {
				left = append(left, sl[i])
			}
			if sl[i] >= sl[mid] {
				right = append(right, sl[i])
			}
		}
		return append(append(QuickSort(left), sl[mid]), QuickSort(right)...) // left(sort) + mid + right(sort)
	}
}

func main() {

	var ar []int
	if ar == nil {
		fmt.Println("sdfg")
	}
	ar = append(ar, 1)
	if ar == nil {
		fmt.Println("sddfghfg")
	}

}

//a := []int{-1, 2, -3, -4, 5, -6, 7, -8, -9}	//fmt.Println(QuickSort(a))
