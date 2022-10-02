package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main() {
	a := 42
	b := a // copy by value
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)

	// pointer syntax is similar
	var c *int = &a
	*c = 22
	fmt.Println(a, b)

	{
		// pointer arithmetic?
		a := [3]int{1, 2, 3}
		b := &a[0]
		c := &a[1]
		//c = &a[1] - 4 // nope!
		fmt.Printf("%v %p %p\n", a, b, c)
	}

	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	ms.foo = 42 // don't need to fanny about with ->, compiler knows the score
	fmt.Println(ms)

	// Go has reduced the feature set to provide better safety

	myMap := map[string]string{}
	myMap["Bob"] = "is grate"

	anotherMap := myMap
	anotherMap["Fred"] = "bad man"

	// anotherMap and myMap are effectively pointers - they are pointing at the same data
}
