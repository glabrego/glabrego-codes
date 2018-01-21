package main

import "fmt"

func routine(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	routine("direct")

	go routine("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("goingroutining")

	var input string
	fmt.Scanln(&input)
	fmt.Println("Go routinado")
}
