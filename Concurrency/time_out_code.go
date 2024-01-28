// Concurrency Practices : Time out Code
package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	go sendValue(ch)

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Request timed out")
	}
}

func sendValue(ch chan string) {

	// time.Sleep(2 * time.Second) //uncomment this to delay in response time from goroutine on channel
	ch <- "Message"
}
