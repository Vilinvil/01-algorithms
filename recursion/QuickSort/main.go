package main

import "fmt"

// QuickSort ascending(по возрастанию)
func QuickSort(sl []int) []int {
	if len(sl) <= 1 {
		return sl
	} else {
		left := make([]int, 0)
		right := make([]int, 0)
		mid := (len(sl) + 1) / 2

		// Pivot нужен , т.к. иногда append затирает элементы старого слайса для оптимизации,
		// и на его место вписывает элементы, и результат возвращает. В таком случае sl[mid] даст некоректный результат, т.к. slice sl изменится
		pivot := sl[mid]
		for _, val := range append(sl[:mid], sl[mid+1:]...) { // sl without sl[mid]
			if val < pivot {
				left = append(left, val)
			} else { // val >= pivot here
				right = append(right, val)
			}
		}
		return append(append(QuickSort(left), pivot), QuickSort(right)...) // left(sort) + mid + right(sort)
	}
}

func main() {
	a := []int{-1, 2, -3, -4, 5, -6, 7, -8, -9}
	fmt.Println(QuickSort(a))
}
