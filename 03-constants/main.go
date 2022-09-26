package main

import (
	"fmt"
)

// enumerated constants
const (
	_ = iota // 0 - don't care about this value
	a = iota // 1
	b = iota // 2
	c = iota // 3
)

// customized 'next value' expressions repeat if omitted
const (
	_  = iota
	KB = 1 << (10 * iota) // 1 << (10 * 1) - 1024
	MB                    // 1 << (10 * 2) - 1024 * 1024
	GB                    // 1 << (10 * 3) - 1024 * 1024 * 1024
)

const (
	isAmerican = 1 << iota // 1
	isEuropean             // 2
	isAdmin                // 4
)

func main() {
	const myConst int = 42
	// myConst = 27 - can't do this!
	fmt.Printf("%v, %T\n", myConst, myConst)

	// constants can also be shadowed - but what sane person would do that?
	{
		const myConst float32 = 6.66
		fmt.Printf("%v, %T\n", myConst, myConst)
	}

	// enumerated constants
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)

	enunValue := a
	fmt.Printf("%v\n", enunValue == a)
	fmt.Printf("%v\n", enunValue == b)

	fileSize := 5000000. // 4.77 MB
	fmt.Printf("%.2fMB\n", fileSize/MB)

	fmt.Println(isAmerican)
	fmt.Println(isEuropean)
	fmt.Println(isAdmin)
}
