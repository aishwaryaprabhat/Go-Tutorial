package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float64 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}

	return float64(sum) / float64(len(s.grades))
}

func main() {
	s1 := Student{"Tim", []int{70, 80, 90}, 19}
	fmt.Println(s1.age)
	s1.setAge(21)
	fmt.Println(s1.age)
	fmt.Println(s1.getAverageGrade())
}
