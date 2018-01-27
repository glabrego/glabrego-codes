package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "Olá, Labrego!" }()

	msg := <-messages
	fmt.Println(msg)
}
