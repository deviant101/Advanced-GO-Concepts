package main

import (
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ch <- 1 //goroutine blocked at this point
		wg.Done()
	}()
	wg.Wait()
	close(ch) //not reaches at this point because goroutine is currently blocked at 'ch<-1' line and done is not executed so wait will continue waiting
	<-ch
}
