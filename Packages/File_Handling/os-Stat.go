package main

import (
	"fmt"
	"os"
)

func main() {

	fileInfo, err := os.Stat("../../Anonymous_Go_Routine.go") //can give absolute path also
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Name : ", fileInfo.Name())
	fmt.Println("isDir : ", fileInfo.IsDir())
	fmt.Println("Size : ", fileInfo.Size())
	fmt.Println("Permissions : ", fileInfo.Mode())
	fmt.Println("Time : ", fileInfo.ModTime())
}
