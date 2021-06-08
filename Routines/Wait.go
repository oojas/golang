package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go func() {
		fmt.Println("This is controlled by a waitgroup")
		wg.Done()
	}()
	wg.Wait()
}
