package main

import (
	"encoding/json"
	"fmt"
)

type Ojas struct {
	Name   string `jsom:"Name"`
	Age    int
	Gender string
}

var riya = []Ojas{
	{Name: "Riya", Age: 21, Gender: "Female"},
	{Name: "Rahul", Age: 22, Gender: "Male"},
	{Name: "Sakshi", Age: 34, Gender: "Male"},
}

func main() {

	data, err := json.MarshalIndent(riya, "", " ")
	if err != nil {
		fmt.Println("Not able to marshal", err)
	}
	fmt.Println(string(data))
}
