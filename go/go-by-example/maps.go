package main

import (
	"fmt"
	"strconv"
)

func main() {
	maps := make(map[string]int)

	maps["value1"] = 12
	maps["value2"] = 13
	maps["value3"] = 14

	print_elements(maps, 1, 3)
	fmt.Println("Number of elements:", len(maps))

	delete(maps, "value1")
	print_elements(maps, 1, 3)

	_, value := maps["value2"]
	fmt.Println("Presence: ", value)
}

func print_elements(m map[string]int, begin, end int) {
	for i := begin; i <= end; i++ {
		key := "value" + strconv.Itoa(i)
		fmt.Println(key, m[key])
	}
}
