package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers(&wg)
	go printLetters(&wg)
	wg.Wait()
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for ch := 'a'; ch <= 'e'; ch++ {
		fmt.Printf("%c\n", ch)
	}
}
