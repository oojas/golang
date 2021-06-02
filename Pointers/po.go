package main

import "fmt"

func main() {
	a := 21
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(*b)
	var c = [3]int{1, 2, 3}
	d := &c
	e := &c[2]
	f := &c[0]
	fmt.Println(c, d)
	fmt.Println(f, e)
	fmt.Printf("%T ", f)
	// C holds the address of the c[0] element because the array elements are stored in contiguos locations.
}
