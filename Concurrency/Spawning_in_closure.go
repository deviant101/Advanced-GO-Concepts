// Concurrency Practice : Spawning Go Routine closure in a loop
// Concurrency Practice : When to use buffered and upbuffered channels

package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		// go func() {		//NOTE: Avoid this closure method to spin-up goroutines
		go func(i int) { //INSTEAD: Pass outer values as argument
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Done")
}
