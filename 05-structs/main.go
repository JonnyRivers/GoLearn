package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// Go has no inheritance - but embedding can be used to extend
type Animal struct {
	// Name has two tags, required and max:"100"
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal   // composition extends the Bird type to include the Animal type fields
	SpeedKPH float32
	CanFly   bool
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])

	// anonymous structs
	someGuy := struct{ name string }{name: "Bob"}
	fmt.Println(someGuy)
	fmt.Println(someGuy.name)

	// structs are value types
	anotherGuy := someGuy
	anotherGuy.name = "Fred"
	fmt.Println(someGuy.name)
	fmt.Println(anotherGuy.name)

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b.Name)
	fmt.Println(b.Animal)

	// access tags
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName(("Name"))
	fmt.Println(field.Tag)
}
