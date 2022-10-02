package main

import (
	"fmt"
)

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// compound interface is similar syntax to struct composition
type WriterCloser interface {
	Writer
	Closer
}

type ConsoleWriter struct{}

// no explicit interface opt-in here, which allows user-defined interfaces that map to existing types!

// this also encourages interface segregation

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func main() {
	cw := ConsoleWriter{}
	cw.Write([]byte("Hello Go!"))

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	// empty interface (like Object in C# but simpler)
	// everything implements the empty interface
	var myObj interface{} = ConsoleWriter{}
	// but it's useless, so need to be cast
	r, ok := myObj.(Writer)
	if ok {
		r.Write([]byte("Hello Go!"))
	}

	// switch on type
	var myBool interface{} = true
	switch myBool.(type) {
	case int:
		fmt.Println("myBool is an integer")
	default:
		fmt.Println("I don't know what myBool is")
	}
}
