package main

import (
	"fmt"
	"time"
)

// select lets you wait on multiple channel operations

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "two"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "one"
	}()

	for i := 0; i < 2; i++ {
		select { // We’ll use select to await both of these values simultaneously, printing each one as it arrives
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}
