package main

import "fmt"

type Point struct {
	x        int32
	y        int32
	isOnGrid bool
}

type Circle struct {
	radius float64
	center *Point
}

func main() {
	var p1 Point = Point{1, 2, true}
	var p2 Point = Point{-1, -2, true}

	fmt.Println(p1.x, p2)
	p1.x = 7

	circlePoint := Point{3, 4, true}
	c1 := Circle{22, &circlePoint}
	fmt.Println(c1.center.x)

}
