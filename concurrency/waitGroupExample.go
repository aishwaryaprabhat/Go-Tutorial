package main

import (
	"fmt"
	"sync"
)

func main() {
	var data int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Running addition go routine")
		data++
	}()

	wg.Wait()
	fmt.Printf("Data is %v\n", data)

	fmt.Println("Main go routine done!")

}
