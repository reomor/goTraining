package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			i, more := <-jobs
			if more {
				fmt.Println("got: ", i)
			} else {
				fmt.Println("empty")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
	}
	close(jobs)
	<-done
}
