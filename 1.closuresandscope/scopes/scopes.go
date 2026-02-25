package main

// *******************Block Scope ******************************************

// func main() {
// 	var x int = 12

// 	if true {
// 		var x int = 11
// 		var y int = 12
// 		fmt.Println("Inner x value is :", x)
// 		fmt.Println("Inner y value is :", y)
// 	}
// 	fmt.Println("Outer x value is", x)
// 	//cant acces inner y outsider {}
// 	// fmt.Println("Inner y from outsider {}",y)
// }

// ***************** Function Block *********************************
// func main() {
// 	//cant acces variables from other function
// 	//fmt.Println(a)
// }
// func operation() {
// 	var a int = 12
// 	fmt.Println(a)
// }

//********************* Package level scope ***********************************************
// var globalVar int = 1200

// func main() {
// 	fmt.Println(globalVar)
// 	globalVar = globalVar + 12
// 	fmt.Println(globalVar)
// }

// **************** shadowing of variables ********************************************

// func main() {
// 	x := 12

// 	{
// 		//Modifying the existing variables
// 		//x = 120

// 		//shadowing the variable by creating new variable
// 		x := 1200
// 		fmt.Println(x)
// 	}
// 	fmt.Println(x)
// }
