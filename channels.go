package main

import "fmt"

func main() {
	message := make(chan string) // Create a new channel with make(chan val-type). Channels are typed by the values they convey

	go func() { message <- "bing" }() // Send a value into a channel using the channel <- syntax.
	// Here we send "ping" to the messages channel we made above, from a new goroutine

	msg := <-message // The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
	fmt.Println(msg)
}
