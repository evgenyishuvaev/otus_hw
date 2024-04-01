package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	msg := "Hello, OTUS!"
	reverseMsg := stringutil.Reverse(msg)
	fmt.Println(reverseMsg)
}
