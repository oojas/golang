package main

import "fmt"

type ojas struct {
	firstName string
	lastName  string
}

func main() {
	g := ojas{}
	g.firstName = "Ojas"
	g.lastName = "Gupta"
	g.greet() // this is how we invoke a method.
	var a counter
	b := a.sum(12, 23)
	fmt.Println(b)
}

type counter int

func (g ojas) greet() { // we pass the parameter beofre the name of the function because we get the access to the type from there.
	fmt.Println(g.firstName)
	fmt.Println(g.lastName)
}
func (a counter) sum(b int, c int) int {
	return b + c
}

// A method is basically a function that works with a known context i.e it can work with any type.
