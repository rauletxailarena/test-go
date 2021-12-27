package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return (0.5 * t.base * t.height)
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

type square struct {
	sideLenght float64
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{
		height: 10,
		base:   10,
	}
	s := square{
		sideLenght: 24,
	}
	printArea(t)
	printArea(s)

}
