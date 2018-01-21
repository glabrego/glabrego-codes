package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (rect rectangle) area() float64 {
	return rect.width * rect.height
}

func (rect rectangle) perimeter() float64 {
	return 2*rect.width + 2*rect.height
}

func (circ circle) area() float64 {
	return math.Pi * circ.radius * circ.radius
}

func (circ circle) perimeter() float64 {
	return 2 * math.Pi * circ.radius
}

func measure(geo geometry) {
	fmt.Println(geo)
	fmt.Println(geo.area())
	fmt.Println(geo.perimeter())
}

func main() {
	rectangle := rectangle{width: 3, height: 4}
	circle := circle{radius: 5}

	measure(rectangle)
	measure(circle)
}
