package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() { // Healthchecker goroutine.
		lastTick := time.Now()
		for now := range time.Tick(time.Second) {
			realSleep := now.Sub(lastTick)
			fmt.Printf("heartbeat: %s\n", realSleep)
			lastTick = now

			// Try to detect the case when some CPU-bound task is stealing our
			// time.
			//
			// Note that it is only makes sense when some runtime.Gosched() or
			// other functions are called during that task. In other way it will
			// leave our healthchecker (this) goroutine without CPU time at all.
			if diff := realSleep - time.Second; diff > 0 && diff > time.Second {
				panic("WTF")
			}
		}
	}()
	go func() { // Non-cooperative goroutine.
		for i := 0; ; i++ {
			if i%5000000000 == 0 {
				runtime.Gosched()
			}
		}
	}()

	// Wait for a while.
	<-time.After(15 * time.Second)
}
