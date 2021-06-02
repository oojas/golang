package main

import "fmt"

func main() {
	// in Go the increment operation is not an expression but a statement.

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// for initialsing two values we need to make the increment expression as a statement.
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 { // this is how you initialise two statements at one time.
		fmt.Println(i, j)
	}
	// for range loops
	var s = []int{1, 2, 3, 4}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
