package main

import (
	"fmt"
)

func sayMessage(msg string) {
	fmt.Println(msg)
}

// variables of the same type can save a little space
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted" // does nowt outside as name is passed by value
}

func sayGreetingWithPointers(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
}

// variadic parameters - there can only be one and it must be the last
func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, "The sum is ", result)
}

// return values
func sumReturn(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

// return pointers!?
func sumReturnPointer(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result // auto-promotion to heap!
}

// names return variable
func sumReturnNaming(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

// idiomatic error handling example - with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// methods
type greeter struct {
	greeting string
	name     string
}

// struct method with this by value
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// struct method with this by reference
func (g *greeter) setName(newName string) {
	g.name = newName
}

func main() {
	sayMessage("hello")
	name := "Stacey"
	sayGreeting("Hello", name)
	fmt.Println(name)
	greeting := "Hello"
	sayGreetingWithPointers(&greeting, &name)
	fmt.Println(name)
	sum("Blah!", 1, 2, 3, 4, 5)
	fmt.Println(sumReturn(1, 2, 3))
	fmt.Println(*sumReturnPointer(1, 2))
	fmt.Println(sumReturnNaming(1))
	r, err := divide(1, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)

	// anonymous functions
	var anonFunc func(int) int = func(i int) int {
		fmt.Println("Anonymous!", i)
		return i
	}
	anonFunc(3)
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	g.setName("Fred")
	fmt.Println("The new name is:", g.name)
}
