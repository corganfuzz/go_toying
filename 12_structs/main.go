package main

import (
	"fmt"
	"strcov"
)

// Person is a person
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Grreting method (value receiver)

func (p Person) greet() {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strcov.Itoa(p.age)
}

func main() {
	//init person using struct
	// person1 := Person{firstName: "Sam", lastName: "Smith", city: "Boston", gender: "f", age: 25}

	// person1 := Person{"Sam", "Smith", "Boston", "f", 25}

	// person1.age++
	// fmt.Println(person1.firstName)
	// fmt.Println(person1)

	fmt.Println(person1.greet)

}
