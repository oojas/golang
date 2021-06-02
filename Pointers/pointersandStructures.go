package main

import "fmt"

type myStruct struct {
	foo  int
	name string
}

func main() {
	var a myStruct
	a = myStruct{foo: 21, name: "OJas"}
	fmt.Println(a)
	var b *myStruct
	b = &myStruct{foo: 234, name: "Rahul"}
	fmt.Println(b) // the output is &{234 Rahul} which means that b is holding the address of the struct which has
	// variables foo and name that holds 234 , Rahul values respectively
}
