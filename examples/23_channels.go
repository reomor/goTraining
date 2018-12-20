package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()
	// blocks until receiver and sender are not ready
	msg := <-messages
	fmt.Println(msg)
}
