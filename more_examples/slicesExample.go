package main

import "fmt"

func main() {
	//slices solve problems with arrays like fixed size
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3]

	fmt.Println(cap(s))

	//add elements to slice
	var a []int = []int{1, 2, 3, 4}
	//adds and returns a new slice
	b := append(a, 10)
	fmt.Println(a, b)

	//using make()
	c := make([]int, 5)
	fmt.Printf("%T \n", c)

}
