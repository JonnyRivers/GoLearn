package main

import (
	"fmt"
	"strconv"
)

// package scope
var privateVar int = 66
var PublicVar int = 44

func main() {
	// decalre i of type in
	var i int
	// assign it
	i = 42
	fmt.Printf("%v, %T\n", i, i)

	// delcare and assign in one go
	var j int = 27
	fmt.Printf("%v, %T\n", j, j)

	// nice and succinct
	k := 23
	fmt.Printf("%v, %T\n", k, k)

	// var block
	var (
		name string = "Jonny"
		age  int    = 21
	)

	fmt.Printf("%v is %v\n", name, age)

	{
		// shadowing (with block scope)
		var name int = 55
		fmt.Println(name)
	}

	// string conversions are in a separate package
	seven := strconv.Itoa(7)
	fmt.Println(seven)

	// numeric conversions look like built in functions
	real := 66.66
	converted := int(real)
	println(real, " was converted to ", converted)
}
