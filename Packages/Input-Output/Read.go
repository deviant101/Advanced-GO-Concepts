package main

import (
	"fmt"
	"strings"
)

func main() {

	str := strings.NewReader("Random String")

	buf := make([]byte, 4)
	fmt.Println(str)

	for {
		n, err := str.Read(buf)
		fmt.Println(string(buf[:n]), err)
		if err != nil {
			fmt.Println("Breaking")
			break
		}
	}
	fmt.Println(str)
}
