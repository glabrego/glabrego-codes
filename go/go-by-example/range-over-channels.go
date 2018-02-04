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

	new_queue := make(chan string, 10)
	new_queue <- "Guilherme"

	go func() {
		for elem := range new_queue {
			fmt.Println(elem)
		}
	}()

	new_queue <- "Labrego"
	var input string
	fmt.Scanln(&input)
}
