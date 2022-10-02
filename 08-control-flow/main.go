package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func deferExample() {
	// defer
	fmt.Println("start")
	defer fmt.Println("middle") // this will run before the function exits
	fmt.Println("end")
}

func deferHTTPExample() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // ensure this is closed before we exit this method
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func panicExample() {
	a, b := 1, 0
	ans := a / b // panic happens here
	fmt.Println(ans)
}

func panicAfterDefer() {
	fmt.Println("start")
	defer fmt.Println("middle") // this will run before the function exits
	panic("somethig bad happened")
	fmt.Println("end")
}

func recoverExample() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

func main() {
	// defer, panic, recover
	deferExample()
	deferHTTPExample()

	// there are no exceptions - but we have panic & recover
	//panicExample()
	//panicAfterDefer()

	// libraries tend not to panic but return errors instead

	// recover
	recoverExample()

	// rethrow is handled by panicing the original error
}
