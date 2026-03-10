package main

import (
	"fmt"
	"math"
)

type Measurement interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}
func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func calculation(m Measurement) {
	fmt.Println(m.area())
	fmt.Println(m.perimeter())
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func main() {
	fmt.Println("We are going to learn Interfaces in this Session")
	rect1 := Rectangle{length: 10, breadth: 10}
	calculation(rect1)
	circ1 := Circle{radius: 10}
	calculation(circ1)
}
