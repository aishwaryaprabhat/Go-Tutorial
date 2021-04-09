package main

import "fmt"

func main() {
	//slices solve problems with arrays like fixed size
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[:]

	fmt.Println(s)

}
