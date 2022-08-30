package main

import "fmt"

// Factor calculate factorial of number
func Factor(num int) (int, error) {
	if num < 0 {
		return 0, fmt.Errorf(`function "Factor" use only positive numbers`)
	}
	if num == 1 || num == 0 {
		return 1, nil
	} else {
		tempRes, _ := Factor(num - 1) // Здесь не проверяем на ошибку, т.к. num > 1 и Fact(num-1) не вызовет ошибки
		return num * tempRes, nil
	}
}

func main() {
	res, err := Factor(6)
	if err != nil {
		fmt.Printf("Error in Factor: %v", err)
	} else {
		fmt.Println("Fact:", res)
	}
}
