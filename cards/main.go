package main

import "fmt"

func main(){
	var card = newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}