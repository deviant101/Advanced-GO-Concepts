package main

import (
	"fmt"
	"os"
)

func main() {
	//---------------Reading entire File-----------------//
	fileData, err := os.ReadFile("../../Anonymous_Go_Routine.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Entire File Data : \n", string(fileData))
	fmt.Println("\n")

	//---------------Reading File in Chunks-----------------//
	File, err := os.Open("../../Anonymous_Go_Routine.go")

	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 10)

	for {
		n, err := File.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(buf[:n]))
	}
	File.Close()
}
