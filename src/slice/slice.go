package main

import (
	"fmt"
)

func main() {

	// declare and create
	var aSlice1 []string
	aSlice1 = make([]string, 5)
	fmt.Println(aSlice1)

	// create empty map automatically with 0 length
	aSlice2 := make([]string, 0)
	fmt.Println(aSlice2)

	// add and access an element
	aSlice2 = append(aSlice2, "ten", "eleven")
	fmt.Println(aSlice2)

	// slice literal
	aSlice3 := []string{"one", "two", "three", "four", "five"}
	fmt.Println(aSlice3)

	// Adding 2 slices by using an operator
	aSlice4 := append(aSlice3, aSlice2...)
	fmt.Println(aSlice4)

	// browse over with for range, order is random
	for i, value := range aSlice3 {
		fmt.Printf("element position '%d' maps to value '%s' \n", i, value)
	}

	// length of slice
	fmt.Printf("# elements in map is '%d'\n", len(aSlice3))

	// Get 1st element
	fmt.Println(aSlice3[0])
	//Get last element
	fmt.Println(aSlice3[len(aSlice3)-1])

	// slicing a slice - another slice is returned
	fmt.Println(aSlice3[0:1])
	fmt.Println(aSlice3[0:2])
	fmt.Println(aSlice3[3:])

	// NOTE: that slices have underlying arrays
}
