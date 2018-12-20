package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message: ", msg)
	default:
		fmt.Println("no messages received")
	}
	msg := "message"
	select {
	case messages <- msg:
		fmt.Println("sent: ", msg)
	default:
		fmt.Println("nothing sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received: ", msg)
	case sig := <-signals:
		fmt.Println("received: ", sig)
	default:
		fmt.Println("no activity")
	}
}
