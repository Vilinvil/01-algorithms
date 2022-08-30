package main

import "fmt"

// FindElemInSliceInt Search element in sorted slice of int. Return index
func FindElemInSliceInt(s []int, item int) (int, error) {
	low := 0
	high := len(s) - 1
	for low <= high {
		mid := (low + high) / 2
		switch {
		case s[mid] == item:
			{
				return mid, nil
			}
		case item <= s[mid]:
			{
				high = mid - 1
			}
		case item > s[mid]:
			{
				low = mid + 1
			}
		}
	}
	return 0, fmt.Errorf("%v not found", item)

}

func main() {
	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, i+10)
	}

	index, err := FindElemInSliceInt(a, 15)
	if err != nil {
		fmt.Printf("Error in FindElemInSliceInt: %v", err)
	} else {
		fmt.Printf("Index: %v", index)
	}
}