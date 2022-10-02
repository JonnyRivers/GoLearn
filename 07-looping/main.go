package main

import (
	"fmt"
)

func main() {
	// simple initialize; while; post-iteration
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// initialize and post-iterate is an expression, so multiples are different
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i)
	}

	// some sugar for while-only
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// some sugar for no-while
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	// continue is as expected

	// arrays and slices are iterated using range syntax
	numbers := []int{1, 2, 3}
	for i, v := range numbers {
		fmt.Println(i, v)
	}

	// maps
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862495,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}
}
