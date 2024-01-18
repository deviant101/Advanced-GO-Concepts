package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) //sending 1 for counter in wait group because we know that only 1 goroutine will be selected randomly

	ch := make(chan string)

	go One(ch, &wg)
	go Two(ch, &wg)

	select { //will randomly select if both are ready, otherwise selects the one which gets ready first
	case val := <-ch:
		fmt.Println(val)
		break
		fmt.Println("Not Reachable")
	case val := <-ch:
		fmt.Println(val)
		// default: //to enable this, take care of wait group counter
		// 	fmt.Println("Default case make select statement non-blocking")
	}

	wg.Wait()
}

func One(ch chan string, wg *sync.WaitGroup) {
	ch <- "Channel-1"
	wg.Done()
}

func Two(ch chan string, wg *sync.WaitGroup) {
	// time.Sleep(1 * time.Second)				//can add this line to delay any goroutine for deterministic output
	ch <- "Channel-2"
	wg.Done()
}
