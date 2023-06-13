package main

import "fmt"

// common interface for shapes, bot structs must have `getArea()` method to automatically implement the interface
type area interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLen float64
}

func main() {
	t := triangle{base: 2, height: 5}
	s := square{sideLen: 2}
	fmt.Println("Triangle area =", calculateArea(t))
	fmt.Println("Square area = ", calculateArea(s))

}

func calculateArea(a area) float64 {
	return a.getArea()
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLen * s.sideLen
}
