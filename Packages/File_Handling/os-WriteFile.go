package main

import (
	"fmt"
	"os"
)

func main() {

	File, err := os.OpenFile("../file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600) //go doc os O_RDONLY

	if err != nil {
		fmt.Println(err)
	}

	defer File.Close()

	_, err = File.WriteString("Appending to file")
	if err != nil {
		fmt.Println(err)
	}

}
