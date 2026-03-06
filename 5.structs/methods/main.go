package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) fullNames() string {
	return p.firstName + " " + p.lastName
}
func main() {
	person1 := Person{
		firstName: "Shikhar",
		lastName:  "Dhawan",
		age:       35,
	}
	fmt.Println(person1.fullNames())
}
