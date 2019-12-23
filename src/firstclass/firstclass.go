package main

import (
	"fmt"
)

func doMath(aFunction func(uint32, uint32) float64, a uint32, b uint32) {
	result := aFunction(a, b)
	fmt.Println(result)
}

func sub(a uint32, b uint32) float64 {
	return float64(a - b)
}

func sum(a uint32, b uint32) float64 {
	return float64(a + b)
}

func main() {

	doMath(sum, 5, 4)
	doMath(sub, 3, 1)
}
