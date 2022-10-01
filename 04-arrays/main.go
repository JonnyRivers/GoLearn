package main

import (
	"fmt"
)

func main() {
	// [...] says make it fit what is initialized
	grades := [...]int{97, 85, 93} // arrays are contiguous in memory
	fmt.Printf("Grades: %v\n", grades)

	// [N] is used to decalre an array with default data
	var names [2]string
	names[0] = "Bob"
	names[1] = "Celia"
	fmt.Printf("Names: %v\n", names)

	// arrays are mutable (of course)
	names[0] = "Anne"
	fmt.Printf("Names: %v\n", names)

	// length is provided by the len function
	fmt.Printf("len(names): %v\n", len(names))

	// 2d arrays
	identityMatrix := [3][3]float32{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(identityMatrix)
	// array assignment is copy by default
	matrix := identityMatrix
	matrix[0][1] = 2
	fmt.Println(matrix)
	// assignment by value
	matrixCopy := &identityMatrix
	matrixCopy[0][1] = 3
	fmt.Println(identityMatrix)
}
