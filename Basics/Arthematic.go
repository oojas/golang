package main

import (
	"fmt"
)

func main() {
	a := 12
	var b int = 24
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// Bit Operation
	fmt.Println("Bit Manipulation")
	c := 10
	d := 3
	//     1010
	//     0011
	fmt.Println(c & d)  //    0010
	fmt.Println(c | d)  //    1011
	fmt.Println(c ^ d)  //    1001
	fmt.Println(c &^ d) //    0100

	// Complex Numbers
	e := 2 + 3i
	f := 4 + 6i
	fmt.Println(e + f)
	fmt.Println(e - f)
	fmt.Println(e * f)
	fmt.Println(e / f)
	fmt.Printf("%v, %T\n", e, e) // Just as the float numbers are defaultly stored as float64, complex numbers are stored
	// are complex128 ( float64 for real and float64 for imaginary part)

	// Function for creating complex numbers

	x := complex(7, 8)
	fmt.Println(x)
	fmt.Printf("%v , %T\n", x, x)
}
