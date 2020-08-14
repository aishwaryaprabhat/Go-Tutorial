package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "wkdna;djnf;kasdjn",
		"white": "asijdnpaijsdnfaijsdnf",
		"green": "aksdjnc;kajsnd;kajn",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
