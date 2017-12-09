package main

import "fmt"

func main() {
	slice := make([]string, 3)
	fmt.Println("Emp: ", slice)

	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	fmt.Println("Set: ", slice)
	fmt.Println("Get: ", slice[2])

	fmt.Println("Lenght: ", len(slice))

	slice = append(slice, "d")
	slice = append(slice, "e")
	fmt.Println("Append: ", slice)
	fmt.Println("Lenght: ", len(slice))

	scopy := make([]string, len(slice))
	copy(scopy, slice)
	fmt.Println("Copy: ", scopy)

	l := slice[2:5]
	fmt.Println("Slice slices: ", l)

	c := slice[1:]
	fmt.Println("Slice slices part 2: ", c)

	decl := []string{"x", "y", "z"}
	fmt.Println("Decl: ", decl)

	dimension1 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dimension2 := i + 1
		dimension1[i] = make([]int, dimension2)
		for j := 0; j < dimension2; j++ {
			dimension1[i][j] = i + j
		}
	}
	fmt.Println("Slice2D: ", dimension1)
}
