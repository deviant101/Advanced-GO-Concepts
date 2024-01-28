package main

import (
	"fmt"
	"sync"
	"time"
)

func square(num int, wg *sync.WaitGroup) {
	fmt.Println(num * num)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go square(i, &wg)
	}

	wg.Wait()
	fmt.Println("Elapsed Time:", time.Since(start))
}
