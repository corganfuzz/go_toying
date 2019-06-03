package main

import "fmt"

func main() {

	// var name = "Ger"

	//shorthand

	// name := "Ger"
	// size := 1.3
	name, email := "Ger", "ger@ty.com"

	var age int32 = 37
	var isCool = true

	fmt.Println(name, age, email, isCool)
	fmt.Printf("%T\n", age)
}
