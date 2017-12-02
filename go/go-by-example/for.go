package main

import "fmt"

func main() {
	i := 1

	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j <= 9; j++ {
		fmt.Println("loop")
	}

	for {
		fmt.Println("Guilherme")
		break
	}

	for k := 0; k <= 10; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println(k)
	}
}
