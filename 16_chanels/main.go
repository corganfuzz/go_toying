package main

import (
	"fmt"
	"time"
)

func runLoopSend(n int, ch chan int) {
	for i := 0; i < n; i++ {
		ch <- i // sending the value i to a channel
	}
	close(ch) // this is use to close a channel
}

func runLoopReceived(ch chan int) {
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

	go runLoopSend(10, myChannel)

	go runLoopReceived(myChannel)

	time.Sleep(2 * time.Second)
}
