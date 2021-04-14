package main

import "fmt"

func main() {

	ch := make(chan int, 10)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			fmt.Printf("Sent: %v\n", i)
			ch <- i
		}

	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}
