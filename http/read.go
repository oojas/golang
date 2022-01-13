package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var url string = "https://jsonplaceholder.typicode.com/photos"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Cannot get the data", err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR", err)
	}
	fmt.Printf("%s ", b)
}
