// Concurrency Practice : Cleaning up Go routines
// This is faulty code with memory leak in goroutine
package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go leak(&wg)
	wg.Wait()
	// time.Sleep(1 * time.Second)

}

func leak(WG *sync.WaitGroup) {

	ch := make(chan int)

	go func() {
		val := <-ch //Anonymous go routine will remain blocked because no data is send on channel, will cause in memory leak
		fmt.Println("Received ", val)
		WG.Done()
	}()
	// ch <- 10		//this will unblock Anonymous goroutine by sending data on channel
	fmt.Println("Exiting Leak Method")
	WG.Done()
}
