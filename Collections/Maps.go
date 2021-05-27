package main

import "fmt"

func main() {
	m := map[string]int{
		"ojas":  21,
		"Hiren": 22,
		"Varun": 23,
	}
	// This is  how we create maps.
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(m["ojas"])
	// give values anytime since it is not static
	m["Rabhya"] = 121
	fmt.Println(m["Rabhya"])
	fmt.Println(m["ijas"])
	// Whenever we don't want to declare an actual variable and just see some outcome we use _
	// Map also uses call by reference, this means that if I pass the map to some variable and change something in that variable, it will effect the orgiginal one.
	o := m
	delete(m, "ojas")
	fmt.Println(o)
	fmt.Println(m)
	// Now we can also use , ok to get the instance of whether something is present or not.
	_, ok := m["ojas"]
	fmt.Println(ok)
	// , ok is used specially with maps

}
