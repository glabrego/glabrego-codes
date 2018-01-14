package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func print_elem(m map[string]string) {
	for _, elem := range m {
		fmt.Println(elem)
	}
}

func main() {
	fmt.Println(plus(1, 2))

	person := map[string]string{"name": "Guilherme", "last_name": "Labrego"}
	print_elem(person)
}
