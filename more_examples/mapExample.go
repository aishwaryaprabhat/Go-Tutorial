package main

import "fmt"

func main() {
	//maps allow storing key value pairs
	var mp map[string]int = map[string]int{
		"apple":  1,
		"banana": 2,
	}

	//does not keep track of ordering
	fmt.Println(mp)

	//using make()
	mp2 := make(map[string]int)

	//setting/adding values
	mp2["orange"] = 3
	fmt.Println(mp2)

	//deleting elements
	delete(mp, "apple")

	//checking if a value exists
	val, ok := mp["apple"]
	fmt.Println(ok, val)
}
