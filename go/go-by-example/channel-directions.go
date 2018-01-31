package main

import "fmt"

func ping(pings chan<- string, message string) {
	pings <- message
}

func pong(pings <-chan string, pongs chan<- string) {
	message := <-pings
	pongs <- message
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Guilherme Labrego")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
