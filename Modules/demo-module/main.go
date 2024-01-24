package main

import (
	"fmt"

	"github.com/deviant101/crypto/encrypt"
	"github.com/pborman/uuid"
)

func main() {

	uuid := uuid.NewRandom()
	fmt.Println(uuid)
	fmt.Println(encrypt.Nimbus("go-module"))

}
