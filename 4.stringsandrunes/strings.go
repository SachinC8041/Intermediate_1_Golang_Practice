package main

import "fmt"

func main() {

	//strings in go are read only sequence of bytes i.e. uint8 (not characters;)
	var cricGod string = "Sachin"

	//you cant modify strings directly but you can concat and create new string
	cricketGod := cricGod + " Tendulkar"
	fmt.Println("God of Cricket is : ", cricketGod)

	//find the length of strings in GO (it is length of bytes and not characters)
	fmt.Println("The length of cricketGod var is", len(cricketGod))
	demoString := "é"
	//now the length of demoStirng should be 1 by visual analysis but in go since it's length of bytes
	// it will give 2 bytes answer.
	fmt.Println("The length of demoStirng is", len(demoString)) //2

	//Rune : in Go , Rune represent Character and not byte i.e. int32 whereas
	// strings = "xyz" or `xyz`, Rune = 'x' ,

	runeEx := 'a'
	fmt.Println(runeEx) //return in32 value // ans : 97
	stringEx := "a"
	fmt.Println(stringEx) // returns a string //ans : a

	//Indexing : indexing returns bytes not rune
	var name string = "Raghav Juyal"
	fmt.Println(name[1])        //output will be ascii value i.e. 97
	fmt.Printf("%c\n", name[1]) //output will be a rune/character

	//Raw strings
	rawString := `line 1 		line 2		line3 `
	fmt.Println(rawString)

	//comparing strings
	str1 := "Apple"
	str2 := "apple"
	fmt.Println(str1 <= str2)

	//iterating : Gives ASCII Value
	for i := 0; i < len(cricketGod); i++ {
		fmt.Println(cricketGod[i])
	}

	//iterating : Gives Characters coz of Formatted verbs
	for i := 0; i < len(cricketGod); i++ {
		fmt.Printf("%c\n", cricketGod[i])
	}
	fmt.Println()
	for _, v := range cricketGod {
		fmt.Printf("%c\n", v)
	}

	//converting a rune to string
	convertedRuneToString := string(runeEx)
	fmt.Println(convertedRuneToString)
	fmt.Printf("Type of convertedRuneToString is %T\n", convertedRuneToString)
}
