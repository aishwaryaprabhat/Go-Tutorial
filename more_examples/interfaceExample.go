package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int32
	y int32
}

type Circle struct {
	radius float64
	center *Point
}

type Rectangle struct {
	width  float64
	height float64
}

//any struct that implements area can be considered of type shape
type shape interface {
	area() float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	c1 := Circle{5, &Point{0, 0}}
	r1 := Rectangle{2, 6}

	shapes := []shape{c1, r1}
	for _, v := range shapes {
		fmt.Println(v.area())
	}
}
