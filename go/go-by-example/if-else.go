package main

import "fmt"

func main() {
	if 8%2 == 0 {
		fmt.Println("8 is even")
	} else {
		fmt.Println("8 is odd")
	}

	if nome := "Giovana"; nome == "Harry" {
		fmt.Println("The boy who lived")
	} else if nome == "Giovana" {
		fmt.Println("Lovely girl")
	} else {
		fmt.Println("You should be Guilherme")
	}
}
