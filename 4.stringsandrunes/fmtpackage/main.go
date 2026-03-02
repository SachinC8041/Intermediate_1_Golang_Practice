package main

import "fmt"

func main() {
	//Print() dont go to new line by default
	fmt.Print("Hello ")
	fmt.Print("Go Users")
	fmt.Println()

	//Println() goes to new line by default
	fmt.Println("Hello ")
	fmt.Println("Go users")

	//formatted output using Printf()
	name := "Rahul Dravid"
	fmt.Printf("Name is %s \n", name)

	//sprint and sprintln to concatenate
	//sprint does not add space
	s := fmt.Sprint("Hello ", "Java", "user")
	fmt.Println(s)

	//sprintln() will add space in between
	s2 := fmt.Sprintln("Hello", "Java", "Users")
	fmt.Println(s2)

	//taking input using scan
	var name1 string
	var age int
	//Scan()
	//fmt.Scan(&name1, &age)
	fmt.Scanln(&name1, &age)
	fmt.Println("Name", name1, "Age", age)

}
