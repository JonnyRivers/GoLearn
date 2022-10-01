package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862495,
	}
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Texas"])

	agesByPeople := make(map[string]int)
	agesByPeople["Bob"] = 32
	agesByPeople["Mary"] = 44
	fmt.Println(agesByPeople)
	agesByPeople["Mary"] = 43
	fmt.Println(agesByPeople)
	agesByPeople["Fred"] = 22
	delete(agesByPeople, "Mary")
	fmt.Println(agesByPeople)

	// iteration order is unrelated to insertion order

	// check for presence
	pop, ok := statePopulations["Texas"]
	fmt.Println(pop, ok)

	pop, ok = statePopulations["New Jersey"]
	fmt.Println(pop, ok)

	_, ok = statePopulations["California"]
	fmt.Println(ok)
}
