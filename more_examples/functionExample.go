package main

import "fmt"

//same as func add(x int, y int) int{}
func add(x, y int) int {

	return x + y
}

//returning multiple values
func sumAndProduct(x, y int) (int, int) {
	return x + y, x * y
}

//labelled return values
func sumAndProduct2(x, y int) (z1 int, z2 int) {
	z1 = x + y
	z2 = x * y

	return //returns z1 and z2 that have already been labelled above
}

//defer execution of sth till after the function has returned
func sumAndProduct3(x, y int) (z1 int, z2 int) {
	defer fmt.Println(z1, z1%3 == 0)
	z1 = x + y
	z2 = x * y
	fmt.Println("before return")
	return //returns z1 and z2 that have already been labelled above
}

func main() {
	// fmt.Println(add(2, 2))
	fmt.Println(sumAndProduct3(2, 2))
}
