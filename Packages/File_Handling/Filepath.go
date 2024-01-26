package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := filepath.Join("dir_1", "dir_2/../dir_3", "file.txt")
	fmt.Println(path)

	fmt.Println(filepath.Dir(path))
	fmt.Println(filepath.Base(path))
	fmt.Println(filepath.IsAbs(path))
	fmt.Println(filepath.Ext(path))
}
