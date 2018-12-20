package main

import (
	"fmt"
	"time"
)

func worker(doneChannel chan bool) {
	fmt.Print("working ... ")
	time.Sleep(time.Second * 3)
	fmt.Println("done")
	doneChannel <- true
}

func main() {
	doneChannel := make(chan bool, 1)
	go worker(doneChannel)
	<-doneChannel
}
