package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

// OS threads in C#, etc. are heavyweight - they have 1 MB stack and overhead
// Goroutines are a lightweight abstractions on a thread

var waitGroup = sync.WaitGroup{}
var counter = 0
var mutex = sync.RWMutex{}

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
	// closures allow functions to capture data
	msg := "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye" // there's a race condition here
	time.Sleep(100 * time.Millisecond)

	msg = "Hello"
	go func(msg string) { // passs by value to avoid race
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye" // there's no longer a race condition here
	time.Sleep(100 * time.Millisecond)

	// obviously sleeps are terrible, so wait groups?
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func(index int, msg string) {
			fmt.Println(i, msg)
			waitGroup.Done()
		}(i, "Blah")
	}
	waitGroup.Wait()

	// what about a mutex?
	for i := 0; i < 10; i++ {
		waitGroup.Add(2)
		go func() {
			mutex.RLock()
			fmt.Printf("Hello #%v\n", counter)
			mutex.RUnlock()
			waitGroup.Done()
		}()
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()

	// best practice to *not* create goroutines in libraries
	// when creating a goroutine, know how it will end
	// check for race conditions at compile time (go run -race myapp.go)
}
