package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "First value"
	queue <- "Second value"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
