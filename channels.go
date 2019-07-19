package main

import "fmt"

func main() {
	message := make(chan string)

	go func() { message <- "bing" }()

	msg := <-message
	fmt.Println(msg)
}
