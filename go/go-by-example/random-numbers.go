package main

import "time"
import "fmt"
import "math/rand"

func main() {
	puts := fmt.Print
	putsln := fmt.Println

	puts(rand.Intn(100), ",")
	puts(rand.Intn(100))
	putsln()

	putsln(rand.Float64())

	puts((rand.Float64()*5)+5, ",")
	puts((rand.Float64() * 5) + 5)
	putsln()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	puts(r1.Intn(100), ",")
	puts(r1.Intn(100))
	putsln()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	puts(r2.Intn(100), ",")
	puts(r2.Intn(100))
	putsln()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	puts(r3.Intn(100), ",")
	puts(r3.Intn(100))
}
