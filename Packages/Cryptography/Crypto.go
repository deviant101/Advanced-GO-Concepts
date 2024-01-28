package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func hash(str string) string {
	var Hash = md5.Sum([]byte(str))
	return hex.EncodeToString(Hash[:])
}
func main() {

	Pass := "Password_N0123"
	fmt.Println(hash(Pass))
}
