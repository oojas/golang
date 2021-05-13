package main

import (
	"fmt"
	"strconv"
)

// In order to convert int to string we have to use stringconv packages
func main() {

	var i int = 12
	var j float64
	j = float64(i)
	// var i int=12
	// var j string =(string) i
	fmt.Printf("%v , %T\n", i, i)
	fmt.Printf("%v , %T\n", j, j)
	// Converting Strings
	var k int = 42
	var l string = string(rune(k)) // converting int to string using method doesn't gives this answer because it simply looks for the ASCII character.
	// rune recognizes the integer and converts it into ASCII value.
	fmt.Printf("%v , %T\n", k, k)
	fmt.Printf("%v , %T\n", l, l)
	var m int = 23
	var n string = strconv.Itoa(m)
	fmt.Printf("%v , %T\n", m, m)
	fmt.Printf("%v , %T\n", n, n)
	// we cannot convert strings into int
}
