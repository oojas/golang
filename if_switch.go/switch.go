package main

import "fmt"

func main() {
	var val int = 45

	switch val {
	case 1:
		fmt.Println("This is not the value")
	case 2:
		fmt.Println("This is not the value")
	case 3:
		fmt.Println("This is the value")
	default:
		fmt.Println("value not found")
	}
	// we can also combine cases.
	var v int = 21
	switch v {
	case 1, 23, 21:
		fmt.Println("value found")
	case 34:
		fmt.Println("Value not found")
	default:
		fmt.Println("THis is not the best practice")
	}
	// to get the switch statements mroe verbose we can do the following.
	var ojas int = 21
	switch {
	case ojas <= 12:
		fmt.Println("THis is not the case")
	case ojas >= 23:
		fmt.Println("This is not the case")
	default:
		fmt.Println("You did not gave the right input")
	}
}
