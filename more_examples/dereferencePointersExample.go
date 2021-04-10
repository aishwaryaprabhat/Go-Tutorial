package main

import "fmt"

func changeValue(str *string) { //pass me the pointer
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	//reference is the location where the data is stored
	//& gives us the pointer to x
	x := 7
	y := &x //y is equal to the location of x
	fmt.Println(x, y)

	//dereference operator
	//gives us the value at that address
	//change value through the pointer
	*y = 8
	fmt.Println(x, y)

	toChange := "hello"

	fmt.Println(toChange)
	changeValue2(toChange)
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

}
