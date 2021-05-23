package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6} // Normal declaration of slices

	var b = []int{1, 2, 3, 4, 5, 6} // preffered declaration of slices
	fmt.Println(a[0])
	a[1] = 23
	fmt.Println(a[1])
	fmt.Println(b)
	// Unlike arrays slices don't work with call by value. They work with call by reference and work just as same as the arrays.
	// Also there are several other ways to form a slice.
	var c = make([]int, 3)
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
	fmt.Println(b[2:]) // It displays elements from 2 till the end. The left element of the colon is inclusive
	// but the right side of the colon is exclusive.
}
