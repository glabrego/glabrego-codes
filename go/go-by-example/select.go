package main

import (
	"fmt"
	"time"
)

func main() {
	channel_one := make(chan string)
	channel_two := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		channel_one <- "Guilherme"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		channel_one <- "Labrego"
	}()

	for i := 0; i < 2; i++ {
		select {
		case message_one := <-channel_one:
			fmt.Println("Received:", message_one)
		case message_two := <-channel_two:
			fmt.Println("Received:", message_two)
		}
	}
}
