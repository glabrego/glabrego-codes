package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	str := "sha1: Guilherme Labrego"

	hash := sha1.New()

	hash.Write([]byte(str))

	bytestr := hash.Sum(nil)

	fmt.Println(str)
	fmt.Printf("%x\n", bytestr)
}
