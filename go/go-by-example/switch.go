package main

import "fmt"
import "time"

func main() {
	name := "Guilherme"
	today := time.Now()

	switch name {
	case "Harry":
		fmt.Println("The boy who lived")
	case "Giovana":
		fmt.Println("Lovely girl")
	default:
		fmt.Println("You should be Guiherme")
	}

	switch today.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a week day")
	}

	switch {
	case today.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("It's a boolean")
		case int:
			fmt.Println("It's a integer")
		case string:
			fmt.Println("It's a string")
		default:
			fmt.Println("I don't know what you are.")
		}
	}

	whatAmI(1)
	whatAmI(true)
	whatAmI("Guilherme")
	whatAmI(2.3)
}
