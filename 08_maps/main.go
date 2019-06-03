package main

import "fmt"

func main() {

	// Define map

	emails := make(map[string]string)

	// Assign kv

	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete one

	delete(emails, "Bob")
	fmt.Println(emails)

}
