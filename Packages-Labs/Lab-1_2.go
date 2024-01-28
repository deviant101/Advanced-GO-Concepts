package main

import (
	"fmt"
	"os"
)

func main() {
	// your code goes here
	fileStream, _ := os.Open("./temp.txt")
	buff := make([]byte, 50)

	for {
		n, err := fileStream.Read(buff)
		fmt.Println(string(buff[:n]))
		if err != nil {
			break
		}
	}
}
