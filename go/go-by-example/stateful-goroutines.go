package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOperation struct {
	key      int
	response chan int
}

type writeOperation struct {
	key      int
	value    int
	response chan bool
}

func main() {
	var readOperations uint64
	var writeOperations uint64

	reads := make(chan *readOperation)
	writes := make(chan *writeOperation)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.response <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.response <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOperation{
					key:      rand.Intn(5),
					response: make(chan int)}
				reads <- read
				<-read.response
				atomic.AddUint64(&readOperations, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOperation{
					key:      rand.Intn(5),
					value:    rand.Intn(100),
					response: make(chan bool)}
				writes <- write
				<-write.response
				atomic.AddUint64(&writeOperations, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOperationsFinal := atomic.LoadUint64(&readOperations)
	fmt.Println("readOperations:", readOperationsFinal)

	writeOperationsFinal := atomic.LoadUint64(&writeOperations)
	fmt.Println("writeOperations:", writeOperationsFinal)
}
