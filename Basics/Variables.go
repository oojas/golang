package main

import (
	"fmt"
)

var j int = 19

// %v is used with Printf statement for declaring the datatype
// %T is used to print the type of the data

// Shadowing means that if you have declared a variable at the package level and you re declare the same variable
// at the function level the function level variable has a higher precedence and hence the value of func level var is printed.
// A variable declared must be used in GO otherwise it gives errors.
func main() {
	//var i int = 12 // This is a way to declare variables in GO. We declare variable the way we speak
	// WE have declared a variable that is of type of int.
	//j := 23 // this type of declaration is used when we are unaware of the type of the data.
	fmt.Println(j)
	var j float64 = 90.0
	fmt.Println(j)
	var n bool = true
	m := false
	fmt.Println(n)
	fmt.Println(m)
}
