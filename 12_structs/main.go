package main

import (
	"fmt"
	"strconv"
)

// Person is a person
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer receiver)

func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//init person using struct
	person1 := Person{firstName: "Sam", lastName: "Smith", city: "Boston", gender: "f", age: 25}

	person2 := Person{"Bob", "Jerk", "Boston", "m", 65}

	// person1.age++
	// fmt.Println(person1.firstName)
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Williams")

	person2.getMarried("Thanos")
	fmt.Println(person2.greet())

}
