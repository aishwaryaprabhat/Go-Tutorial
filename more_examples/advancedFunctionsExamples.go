package main

import "fmt"

//assigning function to a variable
func test() {
	x := test2
	x(2)
}

//assigning function to a variable
func test2(x int) int {
	fmt.Println("Printing from inside test 2", x)
	return 2
}

//anonymous function
func test3() {
	x := func() {

	}
	x()

	y := func() int {
		fmt.Println("inside")
		return 9000
	}()
	fmt.Println(y)
}

//passing a function as an argument and returning a function as an argument
func test4(func(int) int) func(int) int {
	return test2
}

func main() {
	test()
}
