package main

import (
	"fmt"
	"sort"
)

func main() {

	var strs = []string{"Apple", "Around", "Armor", "An"}
	sort.Strings(strs)
	fmt.Println(strs)
}
