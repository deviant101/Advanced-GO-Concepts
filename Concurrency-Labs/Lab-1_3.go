package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		fmt.Println("Inside Goroutine")
		wg.Done()
	}()
	wg.Wait()
}
