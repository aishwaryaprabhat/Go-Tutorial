package main

import "fmt"

func main() {

	salary := 10000

	if salary > 13000 {
		fmt.Println("Decent")
	} else if salary > 15000 {
		fmt.Println("Satisfactory")
	} else {
		fmt.Println("Not good enough")
	}

}
