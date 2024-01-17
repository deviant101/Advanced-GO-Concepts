package main

import (
	"fmt"
	"sync"
)

func main() {
	// var channel chan string
	channel := make(chan string)

	var wg sync.WaitGroup //wait group to see the execution of routines
	wg.Add(2)

	go sell(channel, &wg)
	go buy(channel, &wg)
	wg.Wait()

}

func buy(channel chan string, wg *sync.WaitGroup) {

	fmt.Println("Waiting for data from channel")
	value := <-channel //waiting untill channel is unbuffered
	fmt.Println("Data received from channel : ", value)
	fmt.Println("Received data from channel")
	wg.Done()

}

func sell(channel chan string, wg *sync.WaitGroup) {

	channel <- "Products" //waiting untill channel is unbuffered
	fmt.Println("Sent data to Channel")
	wg.Done()
}
