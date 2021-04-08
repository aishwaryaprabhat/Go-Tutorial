package main

import "fmt"

func main() {
	var num1 int = 8
	var num2 int = 4

	//differnt types wont work
	// var num1 float = 8
	// var num2 int = 4

	//conversion to float<->int
	num2 = int(float64(num1))

	answer := num1 + num2

	fmt.Printf("%d \n", answer)
}
