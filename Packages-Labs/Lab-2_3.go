package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// your code goes here
	in_file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(in_file)

	sum := 0
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += val
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	out_file, err := os.Create("output.txt")
	_, err2 := out_file.WriteString(strconv.Itoa(sum))
	if err != nil {
		fmt.Println(err2)
		return
	}
	defer out_file.Close()
}
