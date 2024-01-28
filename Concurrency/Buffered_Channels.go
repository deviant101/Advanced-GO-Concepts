package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int, 3)

	var wg sync.WaitGroup //wait group to see the execution of routines
	wg.Add(2)

	go sell(channel, &wg)
	wg.Wait()

}

func buy(channel chan int, wg *sync.WaitGroup) {

	fmt.Println("Waiting for data from channel")
	value := <-channel                                  //will remain blocked until channel buffer is empty
	fmt.Println("Data received from channel : ", value) //some data is sent to channel buffer (i.e not empty)
	fmt.Println("Received data from channel")
	wg.Done()

}

func sell(channel chan int, wg *sync.WaitGroup) {

	go buy(channel, wg) //buy routine will get blocked when reaches line(22) in buy routine because until now channel is empty
	fmt.Println("Sending data to channel buffer")
	channel <- 1 //unblocked 'buy' routine because channel buffer is not longer empty
	channel <- 2 //sending data to channel
	channel <- 3
	channel <- 4 //the 'sell' routine will get blocked if 'buy' routine has not received data yet(i.e 'buy' is still blocked)
	fmt.Println("Sent data to Channel")
	wg.Done()
}

//NOTE:
//Sending routine gets blocked when channel buffer tries to exceed
//Receiving routine gets blocked if buffer is empty
