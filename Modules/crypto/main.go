package main

import (
	"fmt"

	"github.com/deviant101/crypto/decrypt"
	"github.com/deviant101/crypto/encrypt"
)

func main() {
	fmt.Println(encrypt.Nimbus("go-module"))
	fmt.Println(decrypt.Nimbus(encrypt.Nimbus("go-module")))
}
