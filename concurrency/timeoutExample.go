package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	// ch2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "ch1"
	}()

	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	ch2 <- "ch2"
	// }()

	select {
	case m1 := <-ch1:
		fmt.Println(m1)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}

}
