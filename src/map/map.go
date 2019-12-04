package main

import (
	"fmt"
)

func main() {

	// declare and create
	var aMap1 map[string]int
	aMap1 = make(map[string]int)
	fmt.Println(aMap1)

	// create empty map automatically
	aMap2 := make(map[string]int)
	fmt.Println(aMap2)

	// add and access an element
	aMap2["ten"] = 10
	fmt.Println(aMap2)

	// map literal
	aMap3 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	fmt.Println(aMap3)

	// comma ok idiom
	var ok bool
	var value int
	value, ok = aMap3["one"]
	fmt.Printf("element 'one' maps to value '%d' with result '%t'\n", value, ok)

	// delete an element and access it
	delete(aMap2, "one")
	_, ok = aMap2["one"]
	fmt.Println(ok)

	// browse over with for range, order is random
	for key, value := range aMap3 {
		fmt.Printf("element '%s' maps to value '%d' \n", key, value)
	}

	fmt.Printf("# elements in map is '%d'\n", len(aMap3))

}
