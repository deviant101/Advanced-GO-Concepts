package main

import (
	"fmt"
	"time"
)

func square(num int) int {
	time.Sleep(1 * time.Millisecond)
	fmt.Println(num * num)
	return (num * num)
}

func main() {

	start := time.Now()
	for i := 1; i <= 1000; i++ {
		square(i)
	}
	fmt.Println("Elapsed Time : ", time.Since(start))
}
