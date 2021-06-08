package main

import (
	"fmt"
	"time"
)

func main() {
	go name()
	var msg string = "OJas"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Ojas"
	time.Sleep(100 * time.Millisecond)
}
func name() {
	fmt.Println("THis is a go routine")
}
