package main

import (
	"fmt"
	"time"
)

func runSomeLoop(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Printing", i)
	}
}

func main() {

	// add "go" to a function to run it concurrently

	go runSomeLoop(20)

	time.Sleep(2 * time.Second)

	fmt.Println("Finished")
}
