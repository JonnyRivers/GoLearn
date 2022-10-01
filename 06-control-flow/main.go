package main

import (
	"fmt"
)

func returnTrue() bool {
	fmt.Println("Returning true")
	return true
}

func main() {
	// if statements
	if true { // no single-line blcoks allowed!
		fmt.Println("This test is true")
	}

	// if statements can have initializer-test pairs
	someMap := make(map[string]string)
	someMap["one"] = "1"
	someMap["two"] = "2"
	if val, ok := someMap["one"]; ok {
		fmt.Println(val)
	}

	number := 50
	guess := 30
	if guess < number {
		fmt.Println("Too low")
	} else if guess > number {
		fmt.Println("Too high")
	} else {
		fmt.Println("Bingo!")
	}

	// logical operators are just like C#, etc.
	// short-circuiting works as expected

	// initializer-test pairs are supported here too
	switch num := 2 + 3; num {
	case 1:
		fmt.Println("one")
		// break is not required
	case 2:
		fmt.Println("two")
	case 3, 5: // rather than rely on pass through
		fmt.Println("three or five")
	default:
		fmt.Println("nope")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // the opposite of break, when requires, which is probably rare
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("something else")
	}
}
