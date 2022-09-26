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

	// slices!
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	fmt.Printf("len(names): %v\n", len(slice))
	fmt.Printf("cap(names): %v\n", cap(slice)) // they have capacity (arrays do not)

	// slice assignment is by value! (unlike with array)
	slice2 := slice
	slice2[0] = 4
	fmt.Println(slice)

	// slice syntax is just like in Python
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // slice of all elements
	c := a[3:]  // slice from 4th element (index 3) to end
	d := a[:6]  // slice first 6 elements
	e := a[3:6] // slice 4th thru 6th elements
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// make
	f := make([]int, 3, 100)
	fmt.Println(f)
	fmt.Printf("len(f): %v\n", len(f))
	fmt.Printf("cap(f): %v\n", cap(f))

	// growth
	f = []int{}
	f = append(f, 1)
	fmt.Println(f)
	fmt.Printf("len(f): %v\n", len(f))
	fmt.Printf("cap(f): %v\n", cap(f))
	f = append(f, 2, 3, 4, 5)
	fmt.Println(f)
	fmt.Printf("len(f): %v\n", len(f))
	fmt.Printf("cap(f): %v\n", cap(f))
}
