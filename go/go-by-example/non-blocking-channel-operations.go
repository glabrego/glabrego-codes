package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case message := <-messages:
		fmt.Println("Received:", message)
	default:
		fmt.Println("No message received")
	}

	message := "Hello, everybody!"
	select {
	case messages <- message:
		fmt.Println("Received:", message)
	default:
		fmt.Println("No message sent")
	}

	select {
	case message := <-messages:
		fmt.Println("Received:", message)
	case signal := <-signals:
		fmt.Println("Signals:", signal)
	default:
		fmt.Println("No activity :c")
	}
}
