package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "Guilherme"
	messages <- "Labrego"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
