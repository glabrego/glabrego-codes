package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers list:", numbers)

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum numbers:", sum)

	for index, num := range numbers {
		if num == 4 {
			fmt.Println("Index:", index)
		}
	}

	keyvalue := map[string]string{"name": "Guilherme", "city": "SP"}
	for i, s := range keyvalue {
		fmt.Printf("%s -> %s\n", i, s)
	}

	name := "Guilherme"
	for index, unicode := range name {
		fmt.Println(index, unicode)
	}
}
