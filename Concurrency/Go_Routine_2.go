package main

import (
	"fmt"
	"time"
)

func start() {
	process()
	fmt.Println("In Start")
}

func process() {
	fmt.Println("In Process")
}

func main() {

	start()
	time.Sleep(1 * time.Millisecond)
	// time.Sleep(1 * time.Second)

}
