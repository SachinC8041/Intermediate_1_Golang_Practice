package main

import "fmt"

//Basic example
// func main() {
// 	x := 12

// 	functionVariable := func() {
// 		x++
// 		fmt.Println(x)
// 	}

// 	functionVariable()

// 	x = 30
// 	functionVariable()
// }
//output will be first 13 and then 31.

//**********************************************
//Returning a closure
// func closureEx() func() int {
// 	counter := 0

// 	return func() int {
// 		counter++
// 		return counter
// 	}
// }
// func main() {
// 	f := closureEx()
// 	fmt.Println(f())
// 	fmt.Println(f())
// 	fmt.Println(f())

// 	f2 := closureEx()
// 	fmt.Println(f2())
// }

// Anonymous function , Closure with parameter

func main() {
	substractorOperation := func() func(int) int {
		countdown := 100
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(substractorOperation(10))
	fmt.Println(substractorOperation(20))

}
