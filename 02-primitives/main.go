package main

import (
	"fmt"
)

func main() {
	// bool
	var b bool = true
	fmt.Printf("%v, %T\n", b, b)

	// defaults to false
	var defaultBool bool
	fmt.Printf("%v, %T\n", defaultBool, defaultBool)

	// numerics uint8, uint16, uint32, uint64, "big" package for more
	var n uint = 42
	fmt.Printf("%v, %T\n", n, n)

	// numeric operators work as they do in C. &, |, ^, <<, >> *, /, %, -, +, etc.

	// for real numbers we have 32-bit and 64-bit precision
	var f float32 = 10.1
	fmt.Printf("%v, %T\n", f, f)

	// complex numbers
	var c complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", c, c)
	re := real(c)
	im := imag(c)
	fmt.Printf("%v, %T\n", re, re)
	fmt.Printf("%v, %T\n", im, im)
	c = complex(5, 12)
	fmt.Printf("%v, %T\n", c, c)

	// strings
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])

	// string are immutable (like in C#)

	// strings can be converted to byte arrays
	by := []byte(s)
	fmt.Printf("%v, %T\n", by, by)

	// UTF32 is supported with runes
}
