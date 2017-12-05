package main

import "fmt"
import "strconv"

func main() {
	var onedimensional [5]int
	fmt.Println("one-dimensional:", onedimensional)

	var multidimensional [2][3]string
	fmt.Println("multi-dimensional:", multidimensional)

	initial := [5]int{1, 2, 3, 4, 5}
	fmt.Println("initial:", initial)
	fmt.Println("initial-len:", len(initial))

	tot := len(onedimensional)
	for i := 0; i < tot; i++ {
		onedimensional[i] = i
	}

	fmt.Println("one-dimensional:", onedimensional)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			multidimensional[i][j] = "guilherme" + strconv.Itoa(i) + strconv.Itoa(j)
		}
	}
	fmt.Println("multi-dimensional:", multidimensional)
}
