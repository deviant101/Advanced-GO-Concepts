package main

import (
	"fmt"
	"time"
)

func square(num int) {

	time.Sleep(1 * time.Millisecond)
	fmt.Println(num * num)
}

func main() {

	start := time.Now()
	for i := 1; i <= 1000; i++ {
		go square(i)
	}
	time.Sleep(2 * time.Millisecond)
	fmt.Println("Elapsed Time : ", time.Since(start))
}
