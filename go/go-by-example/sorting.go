package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"b", "a", "e", "c", "d"}
	sort.Strings(strings)
	fmt.Println("Strings:", strings)

	integers := []int{2, 1, 5, 7, 10, 20, 3}
	sort.Ints(integers)
	fmt.Println("Integers:", integers)

	sorted := sort.IntsAreSorted(integers)
	fmt.Println("Sorted:", sorted)
}
