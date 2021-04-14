package main

import "fmt"

//can only send messages
func sendMessage(ch1 chan<- string) {
	ch1 <- "Hello from the other side"
}

//can only receive on ch1 and can only send on ch2
func relayMessage(ch1 <-chan string, ch2 chan<- string) {
	m := <-ch1
	ch2 <- m
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendMessage(ch1)
	go relayMessage(ch1, ch2)
	m := <-ch2
	fmt.Println(m)
}
