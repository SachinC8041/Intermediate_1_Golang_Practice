package main

import "fmt"

func main() {

	// -------------------------
	// General Formatting Verbs
	// -------------------------
	value := 42
	fmt.Printf("%%v : %v\n", value)
	fmt.Printf("%%#v : %#v\n", value)
	fmt.Printf("%%T : %T\n", value)
	fmt.Printf("%%%% : %%\n")

	fmt.Println()

	// -------------------------
	// Integer Formatting
	// -------------------------
	num := 255

	fmt.Printf("%%b (binary) : %b\n", num)
	fmt.Printf("%%d (decimal) : %d\n", num)
	fmt.Printf("%%+d (show sign) : %+d\n", num)
	fmt.Printf("%%o (octal) : %o\n", num)
	fmt.Printf("%%O (octal with 0o) : %O\n", num)
	fmt.Printf("%%x (hex lowercase) : %x\n", num)
	fmt.Printf("%%X (hex uppercase) : %X\n", num)
	fmt.Printf("%%#x (hex with 0x) : %#x\n", num)

	fmt.Println("Padding examples:")
	fmt.Printf("%%4d : %4d\n", num)
	fmt.Printf("%%-4d : %-4d\n", num)
	fmt.Printf("%%04d : %04d\n", num)

	fmt.Println()

	// -------------------------
	// Float Formatting
	// -------------------------
	floatNum := 123.456789

	fmt.Printf("%%f : %f\n", floatNum)
	fmt.Printf("%%.2f : %.2f\n", floatNum)
	fmt.Printf("%%8.2f : %8.2f\n", floatNum)
	fmt.Printf("%%e : %e\n", floatNum)
	fmt.Printf("%%E : %E\n", floatNum)
	fmt.Printf("%%g : %g\n", floatNum)

	fmt.Println()

	// -------------------------
	// String Formatting
	// -------------------------
	str := "Golang"

	fmt.Printf("%%s : %s\n", str)
	fmt.Printf("%%q : %q\n", str)
	fmt.Printf("%%8s : %8s\n", str)
	fmt.Printf("%%-8s : %-8s\n", str)
	fmt.Printf("%%x (hex) : %x\n", str)
	fmt.Printf("%% X (hex with spaces) : % X\n", str)

	fmt.Println()

	// -------------------------
	// Character / Rune
	// -------------------------
	ch := 'A'

	fmt.Printf("%%c : %c\n", ch)
	fmt.Printf("%%q : %q\n", ch)
	fmt.Printf("%%U : %U\n", ch)

	fmt.Println()

	// -------------------------
	// Boolean
	// -------------------------
	flag := true
	fmt.Printf("%%t : %t\n", flag)

	fmt.Println()

	// -------------------------
	// Pointer
	// -------------------------
	ptr := &value
	fmt.Printf("%%p : %p\n", ptr)

	fmt.Println()

	// -------------------------
	// Struct Formatting
	// -------------------------
	type Person struct {
		Name string
		Age  int
	}

	p := Person{"Sachin", 50}

	fmt.Printf("%%v : %v\n", p)
	fmt.Printf("%%+v : %+v\n", p)
	fmt.Printf("%%#v : %#v\n", p)
}
