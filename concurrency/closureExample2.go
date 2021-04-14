package main

import (
	"fmt"
	"sync"
)

//go routines operate on the current value
//to let them operate on a specific value then we need to pass it a specific value as an input

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Value of i is %v\n", i)
		}(i)

	}

	wg.Wait()
	fmt.Println("All done!")
}
