package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)
	fmt.Println(len(ch))

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	// Wait for the goroutine to finish before sending value to the closed channel
	wg.Wait()

	// Attempting to send data to a closed channel will result in a panic
	// ch <- 10

	for {
		data, ok := <-ch
		if !ok {
			fmt.Println("Channel Closed")
			break
		}
		fmt.Println("Received:", data)
	}

	fmt.Println(len(ch))
}
