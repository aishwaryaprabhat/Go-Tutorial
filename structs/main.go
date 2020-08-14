package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
