package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan string, 2)
	timeout := make(chan bool, 1)

	go One(ch, &wg)
	go Two(ch, &wg)

	go func() {
		time.Sleep(3 * time.Second) // Adjust the timeout duration as needed
		timeout <- true
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	select {
	case val := <-ch:
		fmt.Println(val)
	case val := <-ch:
		fmt.Println(val)
	case <-timeout:
		fmt.Println("Timeout reached. Aborting.")
	}
}

func One(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second) // Simulating work
	ch <- "Channel-1"
}

func Two(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(4 * time.Second) // Simulating work
	ch <- "Channel-2"
}
