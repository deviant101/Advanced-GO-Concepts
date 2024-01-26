package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	str := strings.NewReader("String to write/output")
	// _, err := io.Copy(os.Stdout, str)
	if _, err := io.Copy(os.Stdout, str); err != nil {
		log.Fatal(err)
	}
}
