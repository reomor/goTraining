package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pingsChannel := make(chan string, 1)
	pongsChannel := make(chan string, 1)
	ping(pingsChannel, "passed message")
	pong(pingsChannel, pongsChannel)
	fmt.Println(<-pongsChannel)
}
