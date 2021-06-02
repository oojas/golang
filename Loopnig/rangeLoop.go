package main

import "fmt"

func main() {
	m := map[string]int{
		"Ojas":    21,
		"Rahul":   22,
		"Shweta":  23,
		"Radhika": 24,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println()
	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}
}
