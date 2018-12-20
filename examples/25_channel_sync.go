package main

import (
	"fmt"
	"time"
)

func worker(done_channel chan bool) {
	fmt.Print("working ... ")
	time.Sleep(time.Second * 3)
	fmt.Println("done")
	done_channel <- true
}

func main() {
	done_channel := make(chan bool, 1)
	go worker(done_channel)
	<-done_channel
}
