package main

import (
	"fmt"
	"function/read"
	"log"
)

func main() {

	message, err := read.GetMessage()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Output is: ", message)
}
