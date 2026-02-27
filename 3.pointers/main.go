package main

import "fmt"

func main() {
	var num1 int = 12
	var num2 int = 12
	fmt.Println("The ampisant address of num1 is", &num1)
	fmt.Println("The ampisant address of num2 is ", &num2)
	// the ampisant sigh i.e. & is used to get the address of paritucular thing

	var ptr *int = &num2
	//to  store the address of variable num2 in variable ptr we need to use the pointer i.e. Referencing i.e. *int
	//note : there are two ways to use astrick based on the position at which astrick is getting used.
	// *int represent  "pointer type" with int base
	// *ptr represent "operator" ptr variable with pointer type which will return the "value that pointer type variable
	// ptr points to " : "also known as Dereferencing"
	fmt.Println(ptr)
	fmt.Printf("The actual value with address %d is %d\n", ptr, *ptr)

	var input1 int = 1200
	var input1ptr *int = &input1
	fmt.Println("The original value of input is", input1)
	fmt.Println("The address stored at pointer input1ptr is", input1ptr)
	fmt.Println("The value which pointer input1ptr refers to is", *input1ptr)
	//changing the values from pointer
	*input1ptr = 100
	fmt.Println("The modified original value input1 now becomes :", input1)
	demoPtr(input1ptr)
	fmt.Println(input1)

}

func demoPtr(ptr *int) {
	*ptr = *ptr + 10
}
