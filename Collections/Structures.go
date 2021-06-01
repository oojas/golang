package main

import "fmt"

type ojas struct {
	value   int
	name    string
	versace []string
}

func main() {
	obj := ojas{} // this is one way to initialise the struct
	obj.name = "ojas"
	obj.value = 21
	fmt.Println(obj)
	// this is how we declare structures.
	// if we want to print a specific variable from struct then we use dot operator.
	fmt.Println(obj.name)
	obj1 := ojas{ // this is another way to initialise struct.
		value: 21,
		name:  "joas",
		versace: []string{
			"Rina",
			"Avni",
			"Shweta",
		},
	}
	fmt.Println(obj1)

}
