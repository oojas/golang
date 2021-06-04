package main

import "fmt"

func main() {
	fmt.Println("This is the example of the interfaces")
	r := rect{
		length: 21,
		width:  22,
	}
	R := circle{
		radius: 12,
	}
	printArea(r)
	printArea(R)
}

type circle struct {
	radius int
}
type rect struct {
	length int
	width  int
}
type Shape interface {
	area() int
}

func (r circle) area() int {
	return r.radius * r.radius
}
func (x rect) area() int {
	return x.length * x.width
}
func printArea(s Shape) {
	fmt.Println(s)
}
