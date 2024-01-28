package main

import "fmt"

func main() {

	ch := make(chan int, 10)
	fmt.Println("len1", len(ch))

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		fmt.Println("len2", len(ch))
		close(ch)
		// ch <- 10		//panic: sending to already closed channel
	}()
	// close(ch)	//panic: closing already closed channel
	ch <- 10 //will not generate panic if close function is not yet executed inside routine
	ch <- 20
	fmt.Println("len3", len(ch))
	for {
		data, ok := <-ch //receiver can receive data until buffer is not empty
		if !ok {
			fmt.Println("Channel Closed")
			break
		}
		fmt.Println("Received : ", data)
	}
	fmt.Println(len(ch))
}

//Panic Situations:
//Sending data to already closed channel
//Closing an alredy closed channel
