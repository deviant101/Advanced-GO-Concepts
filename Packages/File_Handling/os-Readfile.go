package main

import (
	"fmt"
	"os"
)

func main() {

	fileData, err := os.ReadFile("../../Anonymous_Go_Routine.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(fileData))
}
