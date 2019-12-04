package main

import (
	"fmt"
)

func main() {

	// declare and create
	var anArray1 [5]string
	anArray1[1] = "one"
	anArray1[2] = "two"
	fmt.Println(anArray1)

	// slice literal
	var anArray2 [6]string = [6]string{"one", "two", "three", "four", "five", "six"}
	fmt.Println(anArray2)

	// browse over with for range
	for i, value := range anArray2 {
		fmt.Printf("element position '%d' maps to value '%s' \n", i, value)
	}

	// length of array
	fmt.Printf("# elements in array are '%d'\n", len(anArray2))

	// Get 1st element
	fmt.Println(anArray2[0])
	//Get last element
	fmt.Println(anArray2[len(anArray2)-1])

	// slicing an array - a slice is returned
	fmt.Println(anArray2[0:1])
	fmt.Println(anArray2[0:2])
	fmt.Println(anArray2[3:])
}
