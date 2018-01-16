package main

import "fmt"

func integerSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInteger := integerSequence()

	fmt.Println("First sequence:", nextInteger())
	fmt.Println("First sequence:", nextInteger())
	fmt.Println("First sequence:", nextInteger())

	newIntegers := integerSequence()
	fmt.Println("Second sequence:", newIntegers())

	fmt.Println("First sequence:", nextInteger())
}
