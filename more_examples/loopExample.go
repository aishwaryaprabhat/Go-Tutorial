package main

import "fmt"

func main() {
	//for loops and while loops are both the same in Go

	//similar to while loop
	x := 3
	for x < 5 {
		fmt.Println(x)
		x++
	}

	//Typical for loop
	for x := 0; x <= 5; x += 2 {
		fmt.Println(x * 2)
	}

	//break and continue
	for x := 0; x <= 5; x += 2 {
		break
		continue
	}
}
