package main

import (
	"fmt"
)

func main() {
	var a [2]int
	a[0] = 1
	a[1] = 2
	fmt.Printf("The array values : %v , %v\n", a[0], a[1])
	fmt.Println("The length of the array is", len(a))
	// By default Go doesn't work with call by reference . It works with call by value i.e..
	// it creates copies of arrays and not edits the values of the initial array.
	// if you don't want to copy the data into another array then use & function so that it gives the
	// address of the initial array.
	// So here the concept of pointers comes into play.
	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b)
}
