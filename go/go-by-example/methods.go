package main

import "fmt"

type rectangle struct {
	height, width int
}

func (rect *rectangle) area() int {
	return rect.width * rect.height
}

func (rect rectangle) perimeter() int {
	return 2*rect.width + 2*rect.height
}

func main() {
	rect := rectangle{width: 10, height: 5}
	fmt.Println("Area: ", rect.area())
	fmt.Println("Perimeter: ", rect.perimeter())

	rect_ref := &rect
	fmt.Println("Area: ", rect_ref.area())
	fmt.Println("Perimeter: ", rect_ref.perimeter())
}
