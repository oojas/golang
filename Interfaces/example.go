package main

import "fmt"

func main() {
	f := person{
		firstName: "Ojas",
		lastName:  "Gupta",
		harmones:  "Testosterons",
		sex:       "M",
	}
	f.gender()
	union(f)

}

type person struct {
	firstName string
	lastName  string
	harmones  string
	sex       string
}

func (p person) gender() {
	fmt.Println(p.sex)
}

type orio interface {
	gender()
}

func union(o orio) {
	fmt.Println(o)
}
