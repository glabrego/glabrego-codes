package main

import "os"

func main() {
	panic("a fucking problem")

	_, err := os.Create("tmp/fucking_file")
	if err != nil {
		panic(err)
	}
}
