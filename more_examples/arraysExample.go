package main

import "fmt"

func main() {

	//collection of ordered elements
	//we need to define max size
	var arr [5]int

	//using index
	arr[0] = -1
	fmt.Println(arr)

	//define elements while initializing
	arr2 := [3]int{1, 2, 3}
	fmt.Println(len(arr2))

	//2D array
	arr2D := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2D)

}
