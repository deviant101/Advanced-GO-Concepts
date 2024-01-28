package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("Inside Anonymous Function")
	}()
	time.Sleep(1 * time.Millisecond)
}
