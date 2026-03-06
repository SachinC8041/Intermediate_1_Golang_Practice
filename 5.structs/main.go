package main

import (
	"fmt"
)

type PointerPerson struct {
	firstName string
	lastName  string
	age       int
}
type Person struct {
	name string
	age  int
	pob  string
}

type Employee struct {
	person   Person
	emp_id   int
	emp_dept string
}

// Methods on struct
func (p Person) incrementAge() int {
	p.age = p.age + 1
	return p.age
}

// pointer receiver example
func (p *Person) incrementAgePointerMethod() int {
	p.age++
	return p.age
}

func main() {

	//not initialized structs get defalut values
	var p Person
	fmt.Println(p)

	//way-1 to declare and initialize structs
	var p1 Person
	p1.name = "Rahul Dravid"
	p1.age = 48
	p1.pob = "Bengaluru"
	fmt.Println(p1)

	//way-2 to declare and initialize structs
	p2 := Person{
		name: "Sachin Tendulkar",
		age:  48,
		pob:  "Mumbai",
	}
	fmt.Println(p2)

	//nested structs declaration and initialization way 1
	p3 := Employee{
		person: Person{
			name: "Saurav",
			age:  48,
			pob:  "Kolkata",
		},
		emp_id:   12,
		emp_dept: "Finance",
	}
	fmt.Println(p3)

	//nested structs declaration and initialization way 2
	var p4 Employee
	p4.person.name = "Shardul Thakur"
	p4.person.age = 32
	p4.person.pob = "Pune"
	p4.emp_id = 1221
	p4.emp_dept = "Marketing"
	fmt.Println(p4)

	//Pass by value concept of struct
	person1 := Person{name: "Rahul", age: 12, pob: "Nanded"}
	person2 := Person{name: "XYZ", age: 13, pob: "Pune"}
	person2.name = "Agastya"
	fmt.Println(person1.name)
	fmt.Println(person2.name)

	//Anonymous structs
	anonymousStructs := struct {
		name string
		age  int
	}{
		name: "Sachin Tendulkar",
		age:  12,
	}
	fmt.Println(anonymousStructs)

	//methods on struct example
	fmt.Println(person1.incrementAge())
	fmt.Println(person1.age)

	//pointer receiver example
	fmt.Println(person1.incrementAgePointerMethod())
	fmt.Println(person1.age)

	//Struct Pointers
	pointerPerson := &PointerPerson{firstName: "Sachin", lastName: "Tendulkar", age: 43}
	pointerPerson2 := pointerPerson
	pointerPerson2.age = 31
	fmt.Println("p Age ", pointerPerson.age)
	fmt.Println("p2 Age is", pointerPerson2.age)
}
