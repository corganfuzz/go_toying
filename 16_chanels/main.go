package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{} //Create a global waitgroup

func runLoopSend(n int, ch chan int) {

	defer wg.Done() //ensure wait group counter decrements by one after our function exits

	for i := 0; i < n; i++ {
		ch <- i // sending the value i to a channel "ch"
	}
	close(ch) // this is use to close a channel
}

func runLoopReceived(ch chan int) {

	defer wg.Done() //ensure wait group counter decrements after our function exits

	for {
		i, ok := <-ch // check if the channel is closed
		// if channel is closed value is false, "i gets its value sent from channel
		if !ok {
			break
		}

		fmt.Println("Received value:", i)
	}
}

func main() {

	myChannel := make(chan int)

	wg.Add(2) // Increment the wait group internal counter by 2

	go runLoopSend(10, myChannel)

	go runLoopReceived(myChannel)

	wg.Wait() // Wait till the wait group counter is 0

	time.Sleep(2 * time.Second)
}
