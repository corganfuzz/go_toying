package main

import "fmt"

func main() {

	// Arrays

	var fruitArr [2]string

	//Assign values

	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"

	//Declare and Aasign

	OtherfruitArr := [2]string{"Apple", "Banana"}

	fmt.Println(fruitArr[1])
	fmt.Println(OtherfruitArr)

	//slices

	fruitSlice := []string{"Apple", "Orange", "Cucumber", "Cherry"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
