package main

import (
	"fmt"

	hw02unpackstring "github.com/evgenyishuvaev/otus_hw/hw02_unpack_string/unpack"
)

func main() {
	var inputString string
	fmt.Print("Введите строку для распаковки: ")
	fmt.Scan(&inputString)

	unpackedString, err := hw02unpackstring.Unpack(inputString)
	if err != nil {
		err := fmt.Errorf("fail to unpack: %w", err)
		panic(err)
	}
	fmt.Println(unpackedString)
}
