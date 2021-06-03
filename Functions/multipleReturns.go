package main

import "fmt"

func main() {
	a, b := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum of the numbers is\n", a)
	fmt.Println("The subtraction of the numbers is\n", b)
	a, err := divide(12, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The division of the numbers is ", a)
}

func sum(val ...int) (int, int) {
	result, sub := 0, 0
	for _, v := range val {
		result += v
		sub -= v
	}
	return result, sub
}

// Returning error
func divide(a int, b int) (int, error) {
	if b == 0 {
		return a, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
