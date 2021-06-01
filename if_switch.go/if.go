package main

import "fmt"

func main() {
	if true {
		fmt.Println("My name is Ojas")
	}
	// in GO it is not necessaery to enclose the if statement in brackets and if you have a single statement to run
	// in go you cannot run it by not using curly braces. Curly braces are a must in if statements in go.

	var s int = 12
	var r int = 22
	if s == r {
		fmt.Println("They are in love")
	}
	if r > s {
		fmt.Println("Neha is not in love with him")
	}
	if s < r {
		fmt.Println("She loves someone else")
	}
}
