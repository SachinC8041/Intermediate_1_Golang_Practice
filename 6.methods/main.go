package main

import (
	"fmt"
	"strings"
)

// type Rectangle struct {
// 	length  int
// 	breadth int
// }

// // Method on Struct : Value Receiver type
// // works on copy , doesn't modify original
// func (r Rectangle) calculateArea() int {
// 	return r.length * r.breadth
// }

// //Method on Struct : Pointer Receiver type
// //works on original
// func (r *Rectangle) calculateAreaPointer(scale int) {
// 	r.length *= scale
// 	r.breadth *= scale
// }

// func main() {
// 	rectangle1 := Rectangle{
// 		length:  5,
// 		breadth: 5,
// 	}
// 	fmt.Println("Value receiver area of rectangle:", rectangle1.calculateArea())

// 	rectangle1.calculateAreaPointer(5)
// 	fmt.Println("Pointer receiver area of Rectangle is", rectangle1.calculateArea())

// }

// Methods on Non Struct Types
type DummyInt int

func (d DummyInt) doDouble() int {
	return int(d * 2)
}

type DummyString string

func (s DummyString) doToUpperCase() string {
	return strings.ToUpper(string(s))
}

type Celsius float64

func (c Celsius) toFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

func main() {
	demo := DummyInt(12)
	fmt.Println(demo.doDouble())
	temp := Celsius(100)
	fmt.Println(temp.toFahrenheit()) // Output: 212
	demoString := DummyString("Hello go useR")
	fmt.Println(demoString.doToUpperCase())
}
