package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Opened a ticket at", t)
		}
	}()

	time.Sleep(time.Millisecond * 3100)
	ticker.Stop()
	fmt.Println("Stop opening this fucking tickets!")
}
