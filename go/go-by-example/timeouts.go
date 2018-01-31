package main

import (
	"fmt"
	"time"
)

func main() {
	channel_one := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		channel_one <- "Guilherme"
	}()

	select {
	case response := <-channel_one:
		fmt.Println("Received:", response)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout channel_one")
	}

	channel_two := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		channel_two <- "Labrego"
	}()

	select {
	case response := <-channel_two:
		fmt.Println("Received:", response)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout channel_two")
	}
}
