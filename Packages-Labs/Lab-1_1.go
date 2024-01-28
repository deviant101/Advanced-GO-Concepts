package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("Let us catch up over a cup of coffee")
	// your code goes here
	buff := make([]byte, 5)

	for {
		n, err := reader.Read(buff)
		fmt.Println(buff[:n], err)
		if err != nil {
			break
		}
	}
}
