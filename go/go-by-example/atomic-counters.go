package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter uint64

	// starting 1.000.000 of goroutines to test atomic counters
	for i := 0; i < 1000000; i++ {
		go func() {
			for {
				atomic.AddUint64(&counter, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second * 10)

	opsFinal := atomic.LoadUint64(&counter)
	fmt.Println("counter:", opsFinal)
}
