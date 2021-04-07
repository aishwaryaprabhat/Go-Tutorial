package main

import "fmt"

func main() {
	//explicit declaration
	var num1 uint16 = 260
	num1 = num1 + 5

	//implicit declaration
	var num2 = 2
	num2 = num2 + 5

	//expression assigment
	num3 := 2
	num3 = num3 + 5

	fmt.Println(num1)
	fmt.Printf("%T", num3)

}
