package main

import "fmt"

func main() {

	ans := -2

	//straightforward way of writing switch case
	switch ans {
	case 1:
		fmt.Println("Case 1")
	//multiple cases x or y
	case 2, -2:
		fmt.Println("Case 2")
	default:
		fmt.Println("not a case")
	}

	//switch without variable
	switch {
	case ans > 0:
		fmt.Println("greater than 0")
	case ans < 0:
		fmt.Println("less than 0")
	default:
		fmt.Println("default")
	}

}
