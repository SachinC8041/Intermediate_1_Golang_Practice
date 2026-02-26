package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(3))
	fmt.Println(sumOfDigits(12345))
}

func factorial(a int) int {
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)

}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}
