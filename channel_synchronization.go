package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("Completed")

	done <- true // Send a value to notify that we’re done
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done // Block until we receive a notification from the worker on the channel
}
