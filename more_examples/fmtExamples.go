package main

import "fmt"

//sprintf and printf

func main() {

	//print type
	fmt.Printf("Type %T \n", true)

	//print value
	fmt.Printf("Value %v \n", 10)

	//percent sign
	fmt.Printf("10%% \n")

	//binary octal decimal hexadecimal
	fmt.Printf("Value %b \n", 34335)
	fmt.Printf("Value %o \n", 34335)
	fmt.Printf("Value %d \n", 34335)
	fmt.Printf("Value %x \n", 34335)

	//scientific notation
	fmt.Printf("Value %e \n", 34335.33)
	fmt.Printf("Value %f \n", 34335.33)

	//string
	fmt.Printf("String %s \n", "Aish")
	fmt.Printf("String %q \n", "Aish")

}
