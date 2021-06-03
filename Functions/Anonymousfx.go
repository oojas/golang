package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is an anonymous function")
	}() // this is basicaly indicated to invoke the anonymous function.
	printnum()
}

// We can also pass a function as a type
func printnum() {
	var f func() = func() {
		fmt.Println("This is a function")
	}

	f()
	fmt.Printf("%T", f)
}
