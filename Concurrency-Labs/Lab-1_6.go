package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sell(ch)
	go buy(ch)
	time.Sleep(time.Second * 2) // so main goroutine doesn't exit
}

func sell(ch chan int) {
	ch <- 1
	ch <- 2
	close(ch)
}

func buy(ch chan int) {
	for i := 1; i <= 3; i++ {
		val, ok := <-ch
		fmt.Println(val, ok)
	}

}
