package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	// ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "ch1"
	}()

	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	ch2 <- "ch2"
	// }()

	for {
		select {
		case m1 := <-ch1:
			fmt.Println(m1)
			return
		default:
			// time.Sleep(1 * time.Second)
			fmt.Println("No message yet")
		}
	}

}
