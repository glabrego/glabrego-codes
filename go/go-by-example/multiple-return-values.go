package main

import "fmt"

func numbers() (int, int) {
	return 1, 2
}

func strings() (string, string) {
	return "Guilherme", "Labrego"
}

func main() {
	a, b := numbers()
	fmt.Println(a)
	fmt.Println(b)

	_, c := strings()
	fmt.Println(c)
}
