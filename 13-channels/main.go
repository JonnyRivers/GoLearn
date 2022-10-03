package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var waitGroup = sync.WaitGroup{}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // signal-only channel

func logger() {
	for {
		select { // block on a message being received
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}

func main() {
	{
		// create a channel that recieves messages of a given type
		ch := make(chan int, 50) // channel with 50 integer capacity
		for j := 0; j < 5; j++ {
			waitGroup.Add(2)
			go func(ch <-chan int) { // receive-only channel
				i := <-ch
				fmt.Println(i)
				// ch <- 27 // can't send!
				waitGroup.Done()
			}(ch)
			go func(ch chan<- int) { // send-only channel
				ch <- 42
				waitGroup.Done()
			}(ch)
		}
		waitGroup.Wait()
	}
	{

		go logger()
		logCh <- logEntry{time.Now(), logInfo, "App is starting"}
		for i := 0; i < 5; i++ {
			logCh <- logEntry{time.Now(), logInfo, "Hello"}
		}
		logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
		doneCh <- struct{}{}
	}
}
