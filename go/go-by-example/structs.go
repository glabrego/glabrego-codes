package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Guilherme", 23})
	fmt.Println(person{name: "Giovana", age: 23})
	fmt.Println(person{name: "Agatha"})
	fmt.Println(&person{name: "Raquel", age: 41})

	master := person{"Maciel", 42}
	fmt.Println(master.name)

	master_pointer := &master
	fmt.Println(master_pointer.age)

	master_pointer.age = 43
	fmt.Println(master_pointer.age)
}
