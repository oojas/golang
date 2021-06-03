package main

import "fmt"

func main() {
	defer foo("Hello") // defer variable is basically used to
	sum(1, 2, 3, 4, 5)
	a := multiply(2, 3)
	fmt.Println("The multiplication of the two numbers is ", a)
}
func foo(val string) {
	fmt.Println(val)
}
func multiply(a int, b int) int {
	val := a * b
	return val
}
func sum(values ...int) { // when you have variadic parameters you must pass it as the last argument of the function.
	var result int = 0
	for _, v := range values {
		result += v
	}
	fmt.Printf("The sum of the values is %v\n", result)
}
