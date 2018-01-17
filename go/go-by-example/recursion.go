package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("7! =", factorial(7))
	fmt.Println("11! =", factorial(11))
	fmt.Println("39! =", factorial(39))
}
